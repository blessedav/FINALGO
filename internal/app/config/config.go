package config

import (
	"github.com/joho/godotenv"
)

type Config struct {
	HTTP HttpConfig
}

type HttpConfig struct {
	Addr string `env:"HTTP_CONFIG_ADDR" default:":8080"`
}

func New(filenames ...string) (*Config, error) {
	cfg := new(Config)

	if len(filenames) > 0 {
		if err := godotenv.Load(filenames...); err != nil {
			return nil, err
		}
	}

	// Set default values
	cfg.HTTP.Addr = ":8080"

	return cfg, nil
}
