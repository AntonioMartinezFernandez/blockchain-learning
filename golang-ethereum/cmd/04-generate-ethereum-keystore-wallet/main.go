package main

import (
	"context"
	"fmt"
	"os"

	"github.com/AntonioMartinezFernandez/blockchain-learning/golang-ethereum/config"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	ctx := context.Background()
	cfg, err := config.New[config.Config](ctx)
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	walletKeystore := keystore.NewKeyStore(cfg.KeystoreFolder, keystore.StandardScryptN, keystore.StandardScryptP)

	walletAccount, err := walletKeystore.NewAccount(cfg.WalletPassword)
	if err != nil {
		fmt.Println("Error creating new account:", err)
		os.Exit(1)
	}

	keystoreFilepath := walletAccount.URL.Path

	fmt.Printf("New wallet address: %s - Filepath: %s\n", walletAccount.Address.Hex(), keystoreFilepath)

	fileContentAsBytes, err := os.ReadFile(keystoreFilepath)
	if err != nil {
		fmt.Println("Error reading the keystore file:", err)
		os.Exit(1)
	}

	fmt.Printf("Keystore file content: %s\n", string(fileContentAsBytes))

	decryptedKeystoreKey, err := keystore.DecryptKey(fileContentAsBytes, cfg.WalletPassword)
	if err != nil {
		fmt.Println("Error decrypting the keystore file:", err)
		os.Exit(1)
	}

	privateKeyBytes := crypto.FromECDSA(decryptedKeystoreKey.PrivateKey)
	privateKeyHexString := hexutil.Encode(privateKeyBytes)
	fmt.Printf("Private Key: %s\n", privateKeyHexString)

	publicKeyBytes := crypto.FromECDSAPub(&decryptedKeystoreKey.PrivateKey.PublicKey)
	publicKeyHexString := hexutil.Encode(publicKeyBytes)
	fmt.Printf("Public Key: %s\n", publicKeyHexString)

	addressHexString := decryptedKeystoreKey.Address.Hex()
	fmt.Printf("Address: %s\n", addressHexString)
}
