package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/scrypt"
	"golang.org/x/term"
)

var signTransactionCmd = &cobra.Command{
	Use:   "signTransaction",
	Short: "Sign Transaction using Ed25519",
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

		var encrypted struct {
			Data   string `json:"data"`
			Salt   string `json:"salt"`
			Wallet string `json:"wallet"`
		}
		if err := json.Unmarshal(walletData, &encrypted); err != nil {
			fmt.Println(Red + "❌ Failed to parse the wallet file: " + err.Error() + Reset)
			return
		}

		fmt.Print(Yellow + "🔑 Enter password: " + Reset)
		password, err := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println()
		if err != nil {
			fmt.Println(Red + "❌ Failed to read password." + Reset)
			return
		}

		salt, err := hex.DecodeString(encrypted.Salt)
		if err != nil {
			fmt.Println(Red + "❌ Invalid salt." + Reset)
			return
		}

		// 🔄 Match createWallet: use scrypt
		key, err := scrypt.Key(password, salt, 1<<15, 8, 1, 32)
		if err != nil {
			fmt.Println(Red + "❌ Key derivation failed: " + err.Error() + Reset)
			return
		}

		ciphertext, err := base64.StdEncoding.DecodeString(encrypted.Data)
		if err != nil || len(ciphertext) < 12 {
			fmt.Println(Red + "❌ Invalid ciphertext." + Reset)
			return
		}

		nonce := ciphertext[:12]
		ciphertext = ciphertext[12:]

		block, err := aes.NewCipher(key)
		if err != nil {
			fmt.Println(Red + "❌ AES cipher error: " + err.Error() + Reset)
			return
		}

		aesgcm, err := cipher.NewGCM(block)
		if err != nil {
			fmt.Println(Red + "❌ AES-GCM init error: " + err.Error() + Reset)
			return
		}

		plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			fmt.Println(Red + "❌ Decryption failed. Wrong password?" + Reset)
			return
		}

		var wallet map[string]string
		if err := json.Unmarshal(plaintext, &wallet); err != nil {
			fmt.Println(Red + "❌ Failed to parse decrypted wallet: " + err.Error() + Reset)
			return
		}

		publicKeyHex, ok := wallet["publicKey"]
		privateKeyHex, ok2 := wallet["privateKey"]
		if !ok || !ok2 {
			fmt.Println(Red + "❌ Missing public/private key in wallet." + Reset)
			return
		}

		privateKey, err := hex.DecodeString(privateKeyHex)
		if err != nil || len(privateKey) != ed25519.SeedSize {
			fmt.Println(Red + "❌ Invalid private key." + Reset)
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
