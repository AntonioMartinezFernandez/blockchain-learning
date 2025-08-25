package main

import (
	"context"
	"fmt"
	"os"

	"github.com/AntonioMartinezFernandez/blockchain-learning/golang-ethereum/config"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx := context.Background()
	cfg, err := config.New[config.Config](ctx)
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	client, err := ethclient.DialContext(ctx, cfg.InfuraURL)
	if err != nil {
		fmt.Println("Error creating client:", err)
		os.Exit(1)
	}
	defer client.Close()

	fmt.Println("We are connected to the Ethereum network!")

	block, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		fmt.Println("Error getting latest block:", err)
		os.Exit(1)
	}

	fmt.Printf("Latest block number: %d\n", block.Number().Uint64())
}
