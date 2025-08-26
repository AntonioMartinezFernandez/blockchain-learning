package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/AntonioMartinezFernandez/blockchain-learning/golang-ethereum/config"

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

	client, err := ethclient.DialContext(ctx, cfg.LocalBlockchainURL)
	if err != nil {
		fmt.Println("Error creating client:", err)
		os.Exit(1)
	}
	defer client.Close()

	ethAddress := common.HexToAddress(cfg.LocalBlockchainEthAddress)

	balance, err := client.BalanceAt(ctx, ethAddress, nil)
	if err != nil {
		fmt.Println("Error getting balance:", err)
		os.Exit(1)
	}

	balanceAsEther := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))

	fmt.Printf("The balance of address %s is: %s wei\n", ethAddress.Hex(), balance.String())
	fmt.Printf("The balance of address %s is: %f ether\n", ethAddress.Hex(), balanceAsEther)
}
