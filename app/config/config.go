package config

import (
	"fmt"

	"github.com/vrischmann/envconfig"

	"github.com/katalabut/money-tell-api/app/services/auth"
)

type Config struct {
	HttpPort int `envconfig:"default=8080"`
	Auth     auth.Config
	Mongo    Mongo
}

type (
	Mongo struct {
		Uri      string
		DataBase string
	}
)

func InitConfig(prefix string) (*Config, error) {
	config := &Config{}

	if err := envconfig.InitWithPrefix(config, prefix); err != nil {
		return nil, fmt.Errorf("init config failed: %w", err)
	}

	return config, nil
}
