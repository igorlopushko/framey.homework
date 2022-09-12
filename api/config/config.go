package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v6"
)

var (
	App Config
)

type Config struct {
	LogLevel string `env:"LOG_LEVEL" env-default:"debug"`
}

func init() {
	if err := env.Parse(&App); err != nil {
		fmt.Printf("Could not load config: %v\n", err)
		os.Exit(-1)
	}
}
