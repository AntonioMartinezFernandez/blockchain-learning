package main

import (
	"context"
	"fmt"
	"os"

	"github.com/AntonioMartinezFernandez/blockchain-learning/golang-ethereum/config"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx := context.Background()
	cfg, err := config.New[config.Config](ctx)
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	// Create two accounts
	decryptedKeystore1, decryptedKeystore2 := createAccounts(cfg)
	address1 := decryptedKeystore1.Address
	address2 := decryptedKeystore2.Address

	// --- Alternatively, use existing accounts by specifying their addresses in the config ---
	// address1 := common.HexToAddress(cfg.SepoliaAddress1)
	// address2 := common.HexToAddress(cfg.SepoliaAddress2)

	fmt.Printf("Account 1 Address: %s\n", address1.Hex())
	fmt.Printf("Account 2 Address: %s\n", address2.Hex())

	// Print account balances
	client, err := ethclient.DialContext(ctx, cfg.InfuraTestnetURL)
	if err != nil {
		fmt.Println("Error creating client:", err)
		os.Exit(1)
	}
	defer client.Close()

	balance1, err := client.BalanceAt(ctx, address1, nil)
	if err != nil {
		fmt.Println("Error getting balance:", err)
		os.Exit(1)
	}

	balance2, err := client.BalanceAt(ctx, address2, nil)
	if err != nil {
		fmt.Println("Error getting balance:", err)
		os.Exit(1)
	}

	fmt.Printf("Balance 1: %s wei - Balance 2: %s wei", balance1.String(), balance2.String())
}

func createAccounts(cfg *config.Config) (*keystore.Key, *keystore.Key) {
	walletKeystore := keystore.NewKeyStore(cfg.KeystoreFolder, keystore.StandardScryptN, keystore.StandardScryptP)

	// Create first account
	walletAccount1, err := walletKeystore.NewAccount(cfg.WalletPassword)
	if err != nil {
		fmt.Println("Error creating new account:", err)
		os.Exit(1)
	}

	keystoreFilepath1 := walletAccount1.URL.Path

	fileContentAsBytes1, err := os.ReadFile(keystoreFilepath1)
	if err != nil {
		fmt.Println("Error reading the keystore file:", err)
		os.Exit(1)
	}

	decryptedKeystore1, err := keystore.DecryptKey(fileContentAsBytes1, cfg.WalletPassword)
	if err != nil {
		fmt.Println("Error decrypting the keystore file:", err)
		os.Exit(1)
	}

	// Create second account
	walletAccount2, err := walletKeystore.NewAccount(cfg.WalletPassword)
	if err != nil {
		fmt.Println("Error creating new account:", err)
		os.Exit(1)
	}

	keystoreFilepath2 := walletAccount2.URL.Path

	fileContentAsBytes2, err := os.ReadFile(keystoreFilepath2)
	if err != nil {
		fmt.Println("Error reading the keystore file:", err)
		os.Exit(1)
	}

	decryptedKeystore2, err := keystore.DecryptKey(fileContentAsBytes2, cfg.WalletPassword)
	if err != nil {
		fmt.Println("Error decrypting the keystore file:", err)
		os.Exit(1)
	}

	return decryptedKeystore1, decryptedKeystore2
}
