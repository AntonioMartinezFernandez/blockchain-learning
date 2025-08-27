package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/AntonioMartinezFernandez/blockchain-learning/golang-ethereum/cmd/07-smart-contract-deployment/todo"
	"github.com/AntonioMartinezFernandez/blockchain-learning/golang-ethereum/config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx := context.Background()
	cfg, err := config.New[config.Config](ctx)
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	address := common.HexToAddress(cfg.SepoliaAddress1)

	client, err := ethclient.DialContext(ctx, cfg.InfuraTestnetURL)
	if err != nil {
		fmt.Println("Error creating client:", err)
		os.Exit(1)
	}
	defer client.Close()

	nonce, err := client.PendingNonceAt(ctx, address)
	if err != nil {
		fmt.Println("Error getting nonce:", err)
		os.Exit(1)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		fmt.Println("Error getting gas price:", err)
		os.Exit(1)
	}

	blockchainID, err := client.NetworkID(ctx)
	if err != nil {
		fmt.Println("Error getting network ID:", err)
		os.Exit(1)
	}

	decryptedKeystoreKey, err := getDecryptedKeystoreKey(
		cfg.SepoliaAccount1KeystoreFilepath,
		cfg.WalletPassword,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(decryptedKeystoreKey.PrivateKey, blockchainID)
	if err != nil {
		fmt.Println("Error creating transactor:", err)
		os.Exit(1)
	}
	transactor.GasPrice = gasPrice
	transactor.GasLimit = 3_000_000
	transactor.Nonce = big.NewInt(int64(nonce))

	add, tx, _, err := todo.DeployTodo(transactor, client)
	if err != nil {
		fmt.Println("Error deploying contract:", err)
		os.Exit(1)
	}

	fmt.Printf("Contract deployed! Wait for the transaction %s to be mined...\n", tx.Hash().Hex())
	fmt.Printf("Contract address: %s\n", add.Hex())
	fmt.Printf("Check contract deployment on: https://sepolia.etherscan.io/address/%s\n", add.Hex())
}

func getDecryptedKeystoreKey(keystoreFilepath string, password string) (*keystore.Key, error) {
	fileContentAsBytes, err := os.ReadFile(keystoreFilepath)
	if err != nil {
		return nil, fmt.Errorf("error reading the keystore file: %w", err)
	}

	decryptedKeystoreKey, err := keystore.DecryptKey(fileContentAsBytes, password)
	if err != nil {
		return nil, fmt.Errorf("error decrypting the keystore file: %w", err)
	}

	return decryptedKeystoreKey, nil
}
