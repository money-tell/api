package config

import (
	"fmt"

	"github.com/vrischmann/envconfig"

	"github.com/katalabut/money-tell-api/app/processors/auth"
)

type (
	Config struct {
		HttpPort int `envconfig:"default=8080"`
		Auth     auth.Config
		Postgres *Postgres
	}

	Postgres struct {
		MasterDSN      string `required:"true"`
		SlaveDSN       string `required:"true"`
		MaxIdleClients int    `envconfig:"default=2"`
		MaxOpenClients int    `envconfig:"default=5"`
	}
)

func InitConfig(prefix string) (*Config, error) {
	config := &Config{}

	if err := envconfig.InitWithPrefix(config, prefix); err != nil {
		return nil, fmt.Errorf("init config failed: %w", err)
	}

	return config, nil
}
