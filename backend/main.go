package main

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Transaction struct {
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
	Amount    float64 `json:"amount"`
	Signature string  `json:"signature"`
}

func main() {
	app := fiber.New()

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

		return c.JSON(fiber.Map{
			"message": "Transaction verified",
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
