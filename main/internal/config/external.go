package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

type ExternalConfig struct {
	Host    string        `env:"EXTERNAL_HOST,notEmpty"`
	Port    uint16        `env:"EXTERNAL_REST_PORT,notEmpty"`
	Timeout time.Duration `env:"EXTERNAL_TIMEOUT" envDefault:"5s"`
	CHost   string        `env:"EXTERNAL_HOST_COPY,notEmpty"`
	CPort   uint16        `env:"EXTERNAL_REST_PORT_COPY,notEmpty"`
}

func LoadExternal() (*ExternalConfig, error) {
	cfg := new(ExternalConfig)
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("env.Parse: %w", err)
	}
	return cfg, nil
}

func (c *ExternalConfig) Endpoint() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *ExternalConfig) CopyEndpoint() string {
	return fmt.Sprintf("%s:%d", c.CHost, c.CPort)
}
