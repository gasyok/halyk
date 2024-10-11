package config

import "fmt"

type PgConfig struct {
	Username    string `env:"POSTGRES_USER,notEmpty"`
	Password    string `env:"POSTGRES_PASSWORD,required"`
	HostDefault string `env:"POSTGRES_HOST_DEFAULT,notEmpty"`
	PortDefault uint16 `env:"POSTGRES_PORT_DEFAULT,notEmpty"`
	Database    string `env:"POSTGRES_DB,notEmpty"`
	SSLMode     bool   `env:"POSTGRES_SSL,required"`
}

func (c *PgConfig) ConnString() string {
	connString := func(host string, port uint16) string {
		sslmode := "disable"
		if c.SSLMode {
			sslmode = "require"
		}
		return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
			c.Username,
			c.Password,
			host,
			port,
			c.Database,
			sslmode,
		)
	}
	return connString(c.HostDefault, c.PortDefault)
}
