package config

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	InfuraURL                       string `env:"INFURA_URL"`
	InfuraTestnetURL                string `env:"INFURA_TESTNET_URL"`
	LocalBlockchainURL              string `env:"LOCAL_BLOCKCHAIN_URL"`
	LocalBlockchainEthAddress       string `env:"LOCAL_BLOCKCHAIN_ETH_ADDRESS"`
	KeystoreFolder                  string `env:"KEYSTORE_FOLDER"`
	WalletPassword                  string `env:"WALLET_PASSWORD"`
	SepoliaAddress1                 string `env:"SEPOLIA_ADDRESS_1"`
	SepoliaAddress2                 string `env:"SEPOLIA_ADDRESS_2"`
	SepoliaAccount1KeystoreFilepath string `env:"SEPOLIA_ACCOUNT_1_KEYSTORE_FILEPATH"`
	SepoliaAccount2KeystoreFilepath string `env:"SEPOLIA_ACCOUNT_2_KEYSTORE_FILEPATH"`
}

func New[CONFIG any](ctx context.Context) (*CONFIG, error) {
	godotenv.Load(".env")

	var cfg CONFIG
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
