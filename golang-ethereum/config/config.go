package config

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	InfuraURL         string `env:"INFURA_URL"`
	InfuraTestnetURL  string `env:"INFURA_TESTNET_URL"`
	GanacheURL        string `env:"GANACHE_URL"`
	GanacheEthAddress string `env:"GANACHE_ETH_ADDRESS"`
	KeystoreFolder    string `env:"KEYSTORE_FOLDER"`
	WalletPassword    string `env:"WALLET_PASSWORD"`
	SepoliaAddress1   string `env:"SEPOLIA_ADDRESS_1"`
	SepoliaAddress2   string `env:"SEPOLIA_ADDRESS_2"`
}

func New[CONFIG any](ctx context.Context) (*CONFIG, error) {
	godotenv.Load(".env")

	var cfg CONFIG
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
