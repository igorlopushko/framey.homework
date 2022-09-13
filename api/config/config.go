// Package config is implemented to path configuration to the API library.
package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v6"
)

// An App is a global application configuration instance.
var (
	App Config
)

// Represents API configuration type.
type Config struct {
	LogLevel string `env:"LOG_LEVEL" env-default:"debug"`
}

func init() {
	if err := env.Parse(&App); err != nil {
		fmt.Printf("Could not load config: %v\n", err)
		os.Exit(-1)
	}
}
