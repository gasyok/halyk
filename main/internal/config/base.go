package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type SplitConfig struct {
	Division float32 `env:"DIVISION,notEmpty"`
	Strategy string  `env:"STRATEGY,notEmpty"`
	Port     uint16  `env:"SERVICE_REST_PORT,notEmpty"`
}

func LoadSplit() (*SplitConfig, error) {
	cfg := new(SplitConfig)
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("env.Parse: $w", err)
	}
	fmt.Println("lol ili ne lol")
	switch cfg.Strategy {
	case "hash":
		break
	case "rand":
		break
	default:
		return nil, fmt.Errorf("wrong startegy format")
	}

	return cfg, nil
}
