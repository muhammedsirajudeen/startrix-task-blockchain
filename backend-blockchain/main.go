package main

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type Transaction struct {
	Sender        string  `json:"sender"`
	Recipient     string  `json:"recipient"`
	Amount        float64 `json:"amount"`
	Signature     string  `json:"signature"`
	PreviousBlock string  `json:"previous_block"`
}

// In-memory store
var (
	transactionHistory = []Transaction{}
	accountBalances    = make(map[string]float64)
	mutex              sync.Mutex
)

func main() {
	app := fiber.New()

	// Initialize genesis block
	initGenesisBlock()

	app.Post("/transaction", func(c *fiber.Ctx) error {
		var tx Transaction
		if err := c.BodyParser(&tx); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON",
			})
		}

		valid, err := verifyTransaction(tx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if !valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid signature",
			})
		}

		mutex.Lock()
		defer mutex.Unlock()

		if accountBalances[tx.Sender] < tx.Amount {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Insufficient balance",
			})
		}

		// Set previous block hash
		lastTx := transactionHistory[len(transactionHistory)-1]
		tx.PreviousBlock = hashTransaction(lastTx)

		accountBalances[tx.Sender] -= tx.Amount
		accountBalances[tx.Recipient] += tx.Amount

		transactionHistory = append(transactionHistory, tx)

		return c.JSON(fiber.Map{
			"message": "Transaction verified and recorded",
		})
	})

	app.Get("/balance/:address", func(c *fiber.Ctx) error {
		address := c.Params("address")
		mutex.Lock()
		defer mutex.Unlock()

		balance := accountBalances[address]
		return c.JSON(fiber.Map{
			"address": address,
			"balance": balance,
		})
	})

	app.Get("/transactions", func(c *fiber.Ctx) error {
		mutex.Lock()
		defer mutex.Unlock()

		return c.JSON(transactionHistory)
	})

	app.Get("/transactions/:address", func(c *fiber.Ctx) error {
		address := c.Params("address")
		mutex.Lock()
		defer mutex.Unlock()

		var associatedTransactions []Transaction
		for _, tx := range transactionHistory {
			if tx.Sender == address || tx.Recipient == address {
				associatedTransactions = append(associatedTransactions, tx)
			}
		}

		if associatedTransactions == nil {
			associatedTransactions = []Transaction{}
		}
		return c.JSON(fiber.Map{
			"transactions": associatedTransactions,
		})
	})

	app.Post("/airdrop", func(c *fiber.Ctx) error {
		type AirdropRequest struct {
			Address string `json:"address"`
		}

		var req AirdropRequest
		if err := c.BodyParser(&req); err != nil || req.Address == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body. 'address' field is required.",
			})
		}

		mutex.Lock()
		defer mutex.Unlock()

		accountBalances[req.Address] += 100

		return c.JSON(fiber.Map{
			"message": "Airdropped 100 coins",
			"address": req.Address,
			"balance": accountBalances[req.Address],
		})
	})

	app.Listen(":3000")
}

func verifyTransaction(tx Transaction) (bool, error) {
	pubKeyBytes, err := hex.DecodeString(tx.Sender)
	if err != nil {
		return false, errors.New("invalid sender public key (not hex)")
	}
	if len(pubKeyBytes) != ed25519.PublicKeySize {
		return false, errors.New("invalid public key length (expected 32 bytes)")
	}

	sigBytes, err := base64.StdEncoding.DecodeString(tx.Signature)
	if err != nil {
		return false, errors.New("invalid base64 signature")
	}
	if len(sigBytes) != ed25519.SignatureSize {
		return false, errors.New("invalid signature length (expected 64 bytes)")
	}

	msg := tx.Sender + tx.Recipient + fmt.Sprintf("%.2f", tx.Amount)
	msgHash := sha256.Sum256([]byte(msg))

	if !ed25519.Verify(pubKeyBytes, msgHash[:], sigBytes) {
		return false, nil
	}

	return true, nil
}

func hashTransaction(tx Transaction) string {
	data := tx.Sender + tx.Recipient + fmt.Sprintf("%.2f", tx.Amount) + tx.Signature + tx.PreviousBlock
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func createGenesisBlock() Transaction {
	return Transaction{
		Sender:        "GENESIS",
		Recipient:     "GENESIS",
		Amount:        0,
		Signature:     "",
		PreviousBlock: "",
	}
}

func initGenesisBlock() {
	genesis := createGenesisBlock()
	transactionHistory = append(transactionHistory, genesis)
}
