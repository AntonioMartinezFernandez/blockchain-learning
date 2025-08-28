package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/AntonioMartinezFernandez/blockchain-learning/golang-ethereum/cmd/08-smart-contract-interaction/todo"
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

	clientAddress := common.HexToAddress(cfg.SepoliaAddress1)

	client, err := ethclient.DialContext(ctx, cfg.InfuraTestnetURL)
	if err != nil {
		fmt.Println("Error creating client:", err)
		os.Exit(1)
	}
	defer client.Close()

	nonce, err := client.PendingNonceAt(ctx, clientAddress)
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

	decryptedKeystore, err := getDecryptedKeystore(
		cfg.SepoliaAccount1KeystoreFilepath,
		cfg.WalletPassword,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	contractAddress := common.HexToAddress(cfg.TodoContractAddress)

	todoInstance, err := todo.NewTodo(contractAddress, client)
	if err != nil {
		fmt.Println("Error instantiating contract:", err)
		os.Exit(1)
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(decryptedKeystore.PrivateKey, blockchainID)
	if err != nil {
		fmt.Println("Error creating transactor:", err)
		os.Exit(1)
	}
	transactor.GasPrice = gasPrice
	transactor.GasLimit = 3_000_000
	transactor.Nonce = big.NewInt(int64(nonce))

	tx, err := todoInstance.AddTask(transactor, "A new task from Go!")
	if err != nil {
		fmt.Println("Error adding task:", err)
		os.Exit(1)
	}

	fmt.Printf("Task %s added...\n", tx.Hash().Hex())

	<-time.After(15 * time.Second)

	allTodos, err := todoInstance.GetAllTasks((&bind.CallOpts{Pending: false}))
	if err != nil {
		fmt.Println("Error getting all tasks:", err)
		os.Exit(1)
	}

	fmt.Printf("All tasks:\n%+v\n", allTodos)
}

func getDecryptedKeystore(keystoreFilepath string, password string) (*keystore.Key, error) {
	fileContentAsBytes, err := os.ReadFile(keystoreFilepath)
	if err != nil {
		return nil, fmt.Errorf("error reading the keystore file: %w", err)
	}

	decryptedKeystore, err := keystore.DecryptKey(fileContentAsBytes, password)
	if err != nil {
		return nil, fmt.Errorf("error decrypting the keystore file: %w", err)
	}

	return decryptedKeystore, nil
}
