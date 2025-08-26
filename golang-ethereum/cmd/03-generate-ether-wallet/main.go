package main

import (
	"context"
	"fmt"
	"os"

	"github.com/AntonioMartinezFernandez/blockchain-learning/golang-ethereum/config"

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

	client, err := ethclient.DialContext(ctx, cfg.GanacheURL)
	if err != nil {
		fmt.Println("Error creating client:", err)
		os.Exit(1)
	}
	defer client.Close()

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println("Error generating private key:", err)
		os.Exit(1)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHexString := hexutil.Encode(privateKeyBytes)
	fmt.Printf("Private Key: %s\n", privateKeyHexString)

	publicKey := privateKey.PublicKey
	publicKeyBytes := crypto.FromECDSAPub(&publicKey)
	publicKeyHexString := hexutil.Encode(publicKeyBytes)
	fmt.Printf("Public Key: %s\n", publicKeyHexString)

	addressHexString := crypto.PubkeyToAddress(publicKey).Hex()
	fmt.Printf("Address: %s\n", addressHexString)

	// Verify that the private key can be reconstructed from the hex string
	decodedPrivateKeyHexString, _ := hexutil.Decode(privateKeyHexString)
	reconstructedPrivateKey, _ := crypto.ToECDSA(decodedPrivateKeyHexString)
	fmt.Println("Private keys matches (reconstructed):", reconstructedPrivateKey.Equal(privateKey))
}
