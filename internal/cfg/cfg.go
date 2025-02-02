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
	Region    string `env:"REGION"`
	TableName string `env:"TABLE_NAME"`
	Version   string
}
