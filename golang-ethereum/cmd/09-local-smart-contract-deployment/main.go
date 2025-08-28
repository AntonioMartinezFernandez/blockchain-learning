package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/AntonioMartinezFernandez/blockchain-learning/golang-ethereum/cmd/07-smart-contract-deployment/todo"
	"github.com/AntonioMartinezFernandez/blockchain-learning/golang-ethereum/config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx := context.Background()
	cfg, err := config.New[config.Config](ctx)
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	clientAddress := common.HexToAddress(cfg.LocalBlockchainEthAddress)

	client, err := ethclient.DialContext(ctx, cfg.LocalBlockchainURL)
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

	decodedPrivateKeyBytes, _ := hexutil.Decode(cfg.LocalBlockchainEthAddressPrivateKey)
	privateKey, _ := crypto.ToECDSA(decodedPrivateKeyBytes)

	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, blockchainID)
	if err != nil {
		fmt.Println("Error creating transactor:", err)
		os.Exit(1)
	}
	transactor.GasPrice = gasPrice
	transactor.GasLimit = 3_000_000
	transactor.Nonce = big.NewInt(int64(nonce))

	deployedContractAddress, tx, _, err := todo.DeployTodo(transactor, client)
	if err != nil {
		fmt.Println("Error deploying contract:", err)
		os.Exit(1)
	}

	fmt.Printf("Contract deployed! Wait for the transaction %s to be mined...\n", tx.Hash().Hex())
	fmt.Printf("Contract address: %s\n", deployedContractAddress.Hex())
}
