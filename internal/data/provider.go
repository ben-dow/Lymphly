package data

import (
	"context"
	"errors"
	"lymphly/internal/cfg"
	"strings"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func PutProvider(ctx context.Context, name, tags string, practice *Practice) (*Provider, error) {

	// Derive Provider Id
	providerId := DeriveProviderId(name, practice.PracticeId)

	provider, err := GetProvider(ctx, providerId)
	if err == nil {
		return provider, nil
	} else if !errors.Is(err, ErrProviderNotFound) {
		return nil, err
	}

	// Create Provider
	provider = &Provider{
		ProviderId: providerId,
		Name:       name,
		Tags:       strings.Split(tags, ","),
		PracticeId: practice.PracticeId,
	}

	// Save Provider
	providerRecord := &ProviderRecord{
		PrimaryKey: NewProviderPrimaryKey(providerId),
		Provider:   provider,
	}
	marshaledRecord, _ := attributevalue.MarshalMap(providerRecord)
	_, err = db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &cfg.Cfg().TableName,
		Item:      marshaledRecord,
	})
	if err != nil {
		return nil, err
	}

	return provider, nil
}

var ErrProviderNotFound = errors.New("provider not found")

func GetProvider(ctx context.Context, providerId string) (*Provider, error) {
	primaryKey, _ := attributevalue.MarshalMap(NewProviderPrimaryKey(providerId))

	res, err := db.GetItem(
		ctx,
		&dynamodb.GetItemInput{
			TableName: &cfg.Cfg().TableName,
			Key:       primaryKey,
		},
	)
	if err != nil {
		return nil, err
	}

	if res.Item == nil {
		return nil, ErrProviderNotFound
	}

	out := &ProviderRecord{}
	err = attributevalue.UnmarshalMap(res.Item, out)
	if err != nil {
		return nil, err
	}
	return out.Provider, nil
}
