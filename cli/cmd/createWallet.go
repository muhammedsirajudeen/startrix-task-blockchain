package cmd

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/scrypt"
)

const (
	Blue   = "\033[34m"
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

type Wallet struct {
	Name       string `json:"name"`
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"` // Seed only — safer than full private key
}

// Derive a key using scrypt
func deriveKey(password string, salt []byte) ([]byte, error) {
	return scrypt.Key([]byte(password), salt, 1<<15, 8, 1, 32)
}

// Encrypt JSON data using AES-GCM
func encrypt(data, password string) (string, []byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", nil, err
	}

	key, err := deriveKey(password, salt)
	if err != nil {
		return "", nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", nil, err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(data), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), salt, nil
}

var createWalletCmd = &cobra.Command{
	Use:   "createWallet",
	Short: "Creates a new wallet and saves it as an encrypted JSON file",
	Long:  `This command creates a new Ed25519 wallet and stores it encrypted using a user-provided password.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(Blue, "Enter a name for the new wallet: ", Reset)

		reader := bufio.NewReader(os.Stdin)
		walletName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(Red+"Failed to read input:", err, Reset)
			return
		}

		walletName = strings.TrimSpace(walletName)
		fileName := fmt.Sprintf("%s.json", strings.ReplaceAll(walletName, " ", "_"))

		if _, err := os.Stat(fmt.Sprintf("./wallet/%s", fileName)); err == nil {
			fmt.Println(Red + "A wallet with this name already exists. Please choose a different name." + Reset)
			return
		} else if !os.IsNotExist(err) {
			fmt.Println(Red+"Failed to check if wallet file exists:", err, Reset)
			return
		}

		fmt.Print(Blue, "Enter a password to encrypt your wallet: ", Reset)
		passwordBytes, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(Red+"Failed to read password:", err, Reset)
			return
		}
		password := strings.TrimSpace(passwordBytes)

		fmt.Printf(Green+"Creating wallet with name: %s\n"+Reset, walletName)

		publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			fmt.Println(Red+"Failed to generate keypair:", err)
			return
		}

		wallet := Wallet{
			Name:       walletName,
			PublicKey:  hex.EncodeToString(publicKey),
			PrivateKey: hex.EncodeToString(privateKey.Seed()),
		}

		jsonData, err := json.Marshal(wallet)
		if err != nil {
			fmt.Println(Red+"Failed to marshal wallet:", err, Reset)
			return
		}

		encryptedData, salt, err := encrypt(string(jsonData), password)
		if err != nil {
			fmt.Println(Red+"Failed to encrypt wallet:", err, Reset)
			return
		}

		output := map[string]string{
			"wallet": wallet.Name,
			"salt":   hex.EncodeToString(salt),
			"data":   encryptedData,
		}

		finalJSON, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			fmt.Println(Red+"Failed to marshal encrypted data:", err, Reset)
			return
		}

		err = os.MkdirAll("./wallet", os.ModePerm)
		if err != nil {
			fmt.Println(Red+"Failed to create wallet directory:", err, Reset)
			return
		}

		err = os.WriteFile(fmt.Sprintf("./wallet/%s", fileName), finalJSON, 0600)
		if err != nil {
			fmt.Println(Red+"Failed to write wallet file:", err, Reset)
			return
		}

		fmt.Printf(Green+"Wallet saved securely to ./wallet/%s\n"+Reset, fileName)
	},
}

func init() {
	rootCmd.AddCommand(createWalletCmd)
}
