package data

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/location"
)

var db *dynamodb.Client
var loc *location.Client

func init() {
	cfg, _ := config.LoadDefaultConfig(context.Background())
	db = dynamodb.NewFromConfig(cfg)
	loc = location.NewFromConfig(cfg)
}

func NewProvider(providerName, providerType, practiceId string) error {
	return nil
}

func NewPractice(Name, Address, Phone, Website string) error {
	return nil
}
