package data

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"golang.org/x/crypto/sha3"
)

var db *dynamodb.Client

func init() {
	cfg, _ := config.LoadDefaultConfig(context.Background())
	db = dynamodb.NewFromConfig(cfg)
}

type PrimaryKey struct {
	PartitionKey string `dynamodbav:"pk"`
	SortKey      string `dynamodbav:"sk"`
}

const (
	ProviderPk string = "providers"
)

func NewProviderPrimaryKey(providerId string) *PrimaryKey {
	return &PrimaryKey{
		PartitionKey: ProviderPk,
		SortKey:      providerId,
	}
}

func DeriveProviderId(name, practiceId string) string {
	h := sha3.New256()
	hashInput := fmt.Sprintf("%s-%s", name, practiceId)
	h.Write([]byte(hashInput))
	providerId := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return providerId
}

type Provider struct {
	ProviderId string   `json:"providerId,omitempty"`
	Name       string   `json:"name,omitempty"`
	PracticeId string   `json:"practiceId,omitempty"`
	Tags       []string `json:"tags,omitempty"`
}

type ProviderRecord struct {
	*PrimaryKey
	*Provider
}

const (
	PracticesPk string = "practices"
)

func NewPracticePrimaryKey(practiceId string) *PrimaryKey {
	return &PrimaryKey{
		PartitionKey: PracticesPk,
		SortKey:      practiceId,
	}
}

func DerivePracticeId(name, address string) string {
	h := sha3.New256()
	hashInput := fmt.Sprintf("%s-%s", name, address)
	h.Write([]byte(hashInput))
	practiceId := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return practiceId
}

type Practice struct {
	PracticeId  string   `json:"practiceId,omitempty"`
	Name        string   `json:"name,omitempty"`
	FullAddress string   `json:"fullAddress,omitempty"`
	Lattitude   float64  `json:"lattitude,omitempty"`
	Longitude   float64  `json:"longitude,omitempty"`
	GeoHash     string   `json:"geoHash,omitempty"`
	Phone       string   `json:"phone,omitempty"`
	Website     string   `json:"website,omitempty"`
	State       string   `json:"state,omitempty"`
	StateCode   string   `json:"stateCode,omitempty"`
	Country     string   `json:"country,omitempty"`
	CountryCode string   `json:"countryCode,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

type PracticeRecord struct {
	*PrimaryKey
	*Practice
}
