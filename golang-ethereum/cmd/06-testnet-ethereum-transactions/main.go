package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/AntonioMartinezFernandez/blockchain-learning/golang-ethereum/config"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx := context.Background()
	cfg, err := config.New[config.Config](ctx)
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	address1 := common.HexToAddress(cfg.SepoliaAddress1)
	address2 := common.HexToAddress(cfg.SepoliaAddress2)

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

	fmt.Printf("Balances before transaction\n - address 1: %s wei - address 2: %s wei\n", balance1.String(), balance2.String())

	// Transfer 0.00019 ether from account 1 to account 2
	nonce, err := client.PendingNonceAt(ctx, address1)
	if err != nil {
		fmt.Println("Error getting nonce:", err)
		os.Exit(1)
	}

	// -> 1 ether = 1e18 wei
	// -> 0.00019 ether = 19e13 wei
	value := int64(19e13) // in wei (0.00019 eth)
	transferAmount := big.NewInt(value)
	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		fmt.Println("Error getting gas price:", err)
		os.Exit(1)
	}

	transaction := types.NewTransaction(nonce, address2, transferAmount, gasLimit, gasPrice, nil)

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

	signedTransaction, err := types.SignTx(
		transaction,
		types.NewEIP155Signer(blockchainID),
		decryptedKeystore.PrivateKey,
	)
	if err != nil {
		fmt.Println("Error signing transaction:", err)
		os.Exit(1)
	}

	err = client.SendTransaction(ctx, signedTransaction)
	if err != nil {
		fmt.Println("Error sending transaction:", err)
		os.Exit(1)
	}

	fmt.Printf("Transaction sent: %s\n", signedTransaction.Hash().Hex())

	// Print account balances after the transaction is mined
	// Note: It may take a few seconds for the transaction to be mined, so we retry a few times
	for range 10 {
		balance1, err = client.BalanceAt(ctx, address1, nil)
		if err != nil {
			fmt.Println("Error getting balance:", err)
			os.Exit(1)
		}

		balance2, err = client.BalanceAt(ctx, address2, nil)
		if err != nil {
			fmt.Println("Error getting balance:", err)
			os.Exit(1)
		}

		fmt.Printf("Balances after transaction\n - address 1: %s wei - address 2: %s wei", balance1.String(), balance2.String())

		<-time.After(2 * time.Second)
	}
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
