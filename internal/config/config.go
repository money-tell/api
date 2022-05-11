package config

type Config struct {
	HTTPPort int
}

func InitConfig() (*Config, error) {
	config := &Config{
		HTTPPort: 8080,
	}
	//if err := envconfig.InitWithPrefix(config, prefix); err != nil {
	//	return nil, fmt.Errorf("init config failed: %w", err)
	//}

	return config, nil
}
