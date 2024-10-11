package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type AuctionConfig struct {
	GRPCPort    uint16 `env:"SERVICE_GRPC_PORT,notEmpty"`
	GatewayPort uint16 `env:"SERVICE_GATEWAY_PORT,notEmpty"`

	PgConfig
}

func LoadAuction() (*AuctionConfig, error) {
	cfg := new(AuctionConfig)
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("env.Parse: %w", err)
	}
	return cfg, nil
}
