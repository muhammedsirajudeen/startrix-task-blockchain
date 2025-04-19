package cmd

import (
	"bufio"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const Blue = "\033[34m"
const Reset = "\033[0m"
const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"

type Wallet struct {
	Name       string `json:"name"`
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"` // Seed only — safer than full private key
}

var createWalletCmd = &cobra.Command{
	Use:   "createWallet",
	Short: "Creates a new wallet and saves it as a json file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(Blue, "Enter a name for the new wallet: ", Reset)

		reader := bufio.NewReader(os.Stdin)
		walletName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(Red+"Failed to read input:", err, Reset)
			return
		}

		walletName = walletName[:len(walletName)-1]
		checkfileName := fmt.Sprintf("%s.json", strings.ReplaceAll(walletName, " ", "_"))
		if _, err := os.Stat(fmt.Sprintf("./wallet/%s", checkfileName)); err == nil {
			fmt.Println(Red + "A wallet with this name already exists. Please choose a different name." + Reset)
			return
		} else if !os.IsNotExist(err) {
			fmt.Println(Red+"Failed to check if wallet file exists:", err, Reset)
			return
		}
		fmt.Printf(Green+"Creating wallet with name: %s\n"+Reset, walletName)
		publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			fmt.Println(Red+"Failed to generate keypair:", err)
			return
		}

		fmt.Printf(Green+"Public Key: %x\n", publicKey)
		wallet := Wallet{
			Name:       walletName,
			PublicKey:  hex.EncodeToString(publicKey),
			PrivateKey: hex.EncodeToString(privateKey.Seed()),
		}

		data, err := json.MarshalIndent(wallet, "", "  ")
		if err != nil {
			fmt.Println("Failed to marshal wallet:", err)
			return
		}

		fileName := fmt.Sprintf("%s.json", strings.ReplaceAll(walletName, " ", "_"))
		err = os.MkdirAll("./wallet", os.ModePerm)
		if err != nil {
			fmt.Println(Red+"Failed to create wallet directory:", err)
			return
		}

		err = os.WriteFile(fmt.Sprintf("./wallet/%s", fileName), data, 0600)
		if err != nil {
			fmt.Println(Red+"Failed to write wallet file:", err)
			return
		}

		fmt.Printf(Green+"Wallet saved to %s\n", fileName)
	},
}

func init() {
	rootCmd.AddCommand(createWalletCmd)
}
