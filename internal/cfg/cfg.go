package cfg

import (
	"os"

	"github.com/caarlos0/env/v11"
)

var defaultCfg *Config

// initialization of configuration at package import time
func init() {
	cfg := &Config{
		Version: "1.0",
	}
	err := env.Parse(cfg)
	if err != nil {
		os.Exit(1)
	}

	defaultCfg = cfg
}

// Accessor for Default Configuration
func Cfg() *Config {
	return defaultCfg
}

// Configuration Definition
type Config struct {
	AppName         string `env:"APP_NAME"`
	Environment     string `env:"ENV_NAME"`
	Region          string `env:"REGION"`
	TableName       string `env:"TABLE_NAME"`
	BasePath        string `env:"BASE_PATH"`
	LogLevel        string `env:"LOG_LEVEL"`
	PoolId          string `env:"POOL_ID"`
	ClientId        string `env:"CLIENT_ID"`
	RadarPrivateKey string `env:"RADAR_PRIVATE_KEY"`
	Version         string
}
