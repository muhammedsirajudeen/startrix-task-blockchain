package cmd

import (
	"crypto/aes"
	"crypto/cipher"
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

var showPublicKeyCmd = &cobra.Command{
	Use:   "showPublicKey",
	Short: "Show public key of the selected wallet",
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

		// Ask for the password
		fmt.Print(Yellow + "🔑 Enter password: " + Reset)
		password, err := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println()
		if err != nil {
			fmt.Println(Red + "❌ Failed to read password." + Reset)
			return
		}

		// Decode the salt
		salt, err := hex.DecodeString(encrypted.Salt)
		if err != nil {
			fmt.Println(Red + "❌ Invalid salt." + Reset)
			return
		}

		// 🔄 Match createWallet: use scrypt to derive the key
		key, err := scrypt.Key(password, salt, 1<<15, 8, 1, 32)
		if err != nil {
			fmt.Println(Red + "❌ Key derivation failed: " + err.Error() + Reset)
			return
		}

		// Decode the encrypted data
		ciphertext, err := base64.StdEncoding.DecodeString(encrypted.Data)
		if err != nil || len(ciphertext) < 12 {
			fmt.Println(Red + "❌ Invalid ciphertext." + Reset)
			return
		}

		// Extract the nonce
		nonce := ciphertext[:12]
		ciphertext = ciphertext[12:]

		// Decrypt the ciphertext
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

		// Parse the decrypted wallet JSON
		var wallet map[string]string
		if err := json.Unmarshal(plaintext, &wallet); err != nil {
			fmt.Println(Red + "❌ Failed to parse decrypted wallet: " + err.Error() + Reset)
			return
		}

		// Get the public key
		publicKeyHex, ok := wallet["publicKey"]
		if !ok {
			fmt.Println(Red + "❌ Missing public key in wallet." + Reset)
			return
		}

		fmt.Println(Green + "✅ Public key: " + publicKeyHex + Reset)
	},
}

func init() {
	rootCmd.AddCommand(showPublicKeyCmd)
}
