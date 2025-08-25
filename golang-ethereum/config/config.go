package config

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	InfuraURL string `env:"INFURA_URL"`
}

func New[CONFIG any](ctx context.Context) (*CONFIG, error) {
	godotenv.Load(".env")

	var cfg CONFIG
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
