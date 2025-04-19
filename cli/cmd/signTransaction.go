/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// signTransactionCmd represents the signTransaction command
var signTransactionCmd = &cobra.Command{
	Use:   "signTransaction",
	Short: "Sign Transaction",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		wallets, err := filepath.Glob("./wallet/*.json")
		if err != nil || len(wallets) == 0 {
			fmt.Println(Red + "❌ No wallet files found in ./wallet")
			return
		}

		fmt.Println(Green + "🔐 Select a wallet:")
		for i, file := range wallets {
			fmt.Printf("[%d] %s\n", i+1, file)
		}

		fmt.Print(Yellow + "Enter the number: ")
		var choice int
		fmt.Scanln(&choice)

		if choice < 1 || choice > len(wallets) {
			fmt.Println(Red + "❌ Invalid choice.")
			return
		}

		selectedWallet := wallets[choice-1]
		fmt.Println(Green+"✅ Selected wallet:", selectedWallet)

		// Load the selected wallet JSON file
		walletData, err := os.ReadFile(selectedWallet)
		if err != nil {
			fmt.Println(Red+"❌ Failed to read the wallet file:", err)
			return
		}
		var wallet map[string]string
		err = json.Unmarshal(walletData, &wallet)
		if err != nil {
			fmt.Println(Red+"❌ Failed to parse the wallet file:", err)
			return
		}

		publicKey, ok := wallet["publicKey"]
		if !ok {
			fmt.Println(Red + "❌ publicKey not found in the wallet file.")
			return
		}

		privateKey, ok := wallet["privateKey"]
		if !ok {
			fmt.Println(Red + "❌ privateKey not found in the wallet file.")
			return
		}
		fmt.Println(Green + "✅ Wallet file loaded successfully.")
		fmt.Print(Yellow + "Enter the public key of the recipient: ")
		var recipientPublicKey string
		fmt.Scanln(&recipientPublicKey)

		if recipientPublicKey == "" {
			fmt.Println(Red + "❌ Recipient public key cannot be empty.")
			return
		}

		fmt.Println(Green+"✅ Recipient public key entered:", recipientPublicKey)
		fmt.Print(Yellow + "Enter the amount to send: ")
		var amount float64
		fmt.Scanln(&amount)

		if amount <= 0 {
			fmt.Println(Red + "❌ Amount must be greater than zero.")
			return
		}

		fmt.Println(Green+"✅ Amount entered:", amount)
		// Create the transaction
		transaction := map[string]interface{}{
			"sender":    publicKey,
			"recipient": recipientPublicKey,
			"amount":    amount,
		}

		// Serialize the transaction to JSON
		transactionData, err := json.Marshal(transaction)
		if err != nil {
			fmt.Println(Red+"❌ Failed to serialize the transaction:", err)
			return
		}
		hash := sha256.Sum256(transactionData)

		privateKeyBytes, err := hex.DecodeString(privateKey)
		if err != nil {
			fmt.Println(Red+"❌ Failed to decode private key:", err)
			return
		}

		d := new(big.Int).SetBytes(privateKeyBytes)
		curve := elliptic.P256()
		priv := new(ecdsa.PrivateKey)
		priv.D = d
		priv.PublicKey.Curve = curve
		priv.PublicKey.X, priv.PublicKey.Y = curve.ScalarBaseMult(d.Bytes())

		// Sign the transaction hash
		r, s, err := ecdsa.Sign(rand.Reader, priv, hash[:])
		if err != nil {
			fmt.Println(Red+"❌ Failed to sign transaction:", err)
			return
		}

		signature := append(r.Bytes(), s.Bytes()...)
		signatureStr := base64.StdEncoding.EncodeToString(signature)

		transaction["signature"] = signatureStr

		finalTx, err := json.MarshalIndent(transaction, "", "  ")
		if err != nil {
			fmt.Println(Red+"❌ Failed to marshal final transaction:", err)
			return
		}

		fmt.Println(Green + "\n✅ Transaction signed successfully Paste it in the website")
		fmt.Println(string(finalTx))

	},
}

func init() {
	rootCmd.AddCommand(signTransactionCmd)

}
