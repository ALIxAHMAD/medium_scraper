package env

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string `env:"bot_token"`
}

func Load(path string) error {
	return godotenv.Load(path)
}

func ParseConfig() (*Config, error) {
	config := new(Config)
	if err := env.Parse(config); err != nil {
		return nil, err
	}
	return config, nil
}
