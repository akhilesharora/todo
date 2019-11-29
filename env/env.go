package env

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Host string `env:"HOST" envDefault:"127.0.0.1"`
	Port string `env:"PORT" envDefault:"8000"`
}

func ConfigFromEnv() (Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("failed to parse Config: %+v", err)
	}
	return cfg, nil
}
