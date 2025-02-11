package cfg

import (
	"os"

	"github.com/caarlos0/env/v11"
)

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

var defaultCfg *Config

func Cfg() *Config {
	return defaultCfg
}

type Config struct {
	AppName         string `env:"APP_NAME"`
	Environment     string `env:"ENV_NAME"`
	Region          string `env:"REGION"`
	TableName       string `env:"TABLE_NAME"`
	BasePath        string `env:"BASE_PATH"`
	LogLevel        string `env:"LOG_LEVEL"`
	RadarPrivateKey string `env:"RADAR_PRIVATE_KEY"`
	Version         string
}
