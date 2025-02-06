package cfg

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/caarlos0/env/v11"
)

var ssmClient *ssm.Client

func init() {
	awsCfg, _ := config.LoadDefaultConfig(context.Background())
	ssmClient = ssm.NewFromConfig(awsCfg)

	cfg := &Config{
		Version: "1.0",
	}
	err := env.Parse(cfg)
	if err != nil {
		os.Exit(1)
	}

	res, err := ssmClient.GetParameter(context.Background(), &ssm.GetParameterInput{
		Name:           aws.String(fmt.Sprintf("%s/%s/key/radar/private", cfg.AppName, cfg.Environment)),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		os.Exit(1)
	}

	cfg.RadarPrivateKey = *res.Parameter.Value

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
	RadarPrivateKey string
	Version         string
}
