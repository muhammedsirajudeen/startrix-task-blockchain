package cmd

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var signTransactionCmd = &cobra.Command{
	Use:   "signTransaction",
	Short: "Sign Transaction using Ed25519",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		wallets, err := filepath.Glob("./wallet/*.json")
		if err != nil || len(wallets) == 0 {
			fmt.Println(Red + "❌ No wallet files found in ./wallet" + Reset)
			return
		}

		fmt.Println(Yellow + "🔐 Select a wallet:" + Reset)
		for i, file := range wallets {
			fmt.Printf("[%d] %s\n", i+1, file)
		}

		fmt.Print(Yellow + "Enter the number: " + Reset)
		var choice int
		fmt.Scanln(&choice)

		if choice < 1 || choice > len(wallets) {
			fmt.Println(Red + "❌ Invalid choice." + Reset)
			return
		}

		selectedWallet := wallets[choice-1]
		fmt.Println(Green + "✅ Selected wallet: " + selectedWallet + Reset)

		walletData, err := os.ReadFile(selectedWallet)
		if err != nil {
			fmt.Println(Red + "❌ Failed to read the wallet file: " + err.Error() + Reset)
			return
		}
		var wallet map[string]string
		err = json.Unmarshal(walletData, &wallet)
		if err != nil {
			fmt.Println(Red + "❌ Failed to parse the wallet file: " + err.Error() + Reset)
			return
		}

		publicKeyHex, ok := wallet["publicKey"]
		if !ok {
			fmt.Println(Red + "❌ publicKey not found in the wallet file." + Reset)
			return
		}
		privateKeyHex, ok := wallet["privateKey"]
		if !ok {
			fmt.Println(Red + "❌ privateKey not found in the wallet file." + Reset)
			return
		}

		_, err = hex.DecodeString(publicKeyHex)
		if err != nil {
			fmt.Println(Red + "❌ Invalid public key format: " + err.Error() + Reset)
			return
		}
		privateKey, err := hex.DecodeString(privateKeyHex)
		if err != nil {
			fmt.Println(Red + "❌ Invalid private key format: " + err.Error() + Reset)
			return
		}
		if len(privateKey) != ed25519.SeedSize {
			fmt.Println(Red + "❌ Expected 32-byte private key seed for Ed25519" + Reset)
			return
		}
		privKey := ed25519.NewKeyFromSeed(privateKey)

		fmt.Print(Yellow + "Enter the public key of the recipient: " + Reset)
		var recipient string
		fmt.Scanln(&recipient)
		if recipient == "" {
			fmt.Println(Red + "❌ Recipient public key cannot be empty." + Reset)
			return
		}

		fmt.Print(Yellow + "Enter the amount to send: " + Reset)
		var amount float64
		fmt.Scanln(&amount)
		if amount <= 0 {
			fmt.Println(Red + "❌ Amount must be greater than zero." + Reset)
			return
		}

		transaction := map[string]interface{}{
			"sender":    publicKeyHex,
			"recipient": recipient,
			"amount":    amount,
		}

		message := publicKeyHex + recipient + fmt.Sprintf("%.2f", amount)
		hash := sha256.Sum256([]byte(message))
		signature := ed25519.Sign(privKey, hash[:])

		transaction["signature"] = base64.StdEncoding.EncodeToString(signature)

		finalTx, err := json.MarshalIndent(transaction, "", "  ")
		if err != nil {
			fmt.Println(Red + "❌ Failed to marshal transaction: " + err.Error() + Reset)
			return
		}

		fmt.Println("\n" + Green + "✅ Transaction signed successfully! Paste it in the website." + Reset)
		fmt.Println(string(finalTx))
	},
}

func init() {
	rootCmd.AddCommand(signTransactionCmd)
}
