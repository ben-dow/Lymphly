package data

import (
	"context"
	"lymphly/internal/cfg"
	"lymphly/internal/geo"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
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

func PutPractice(ctx context.Context, id, name, address, phone, website, tags string) error {

	primaryKey := PrimaryKey{
		PartitionKey: "practices",
		SortKey:      id,
	}

	marshaledKey, _ := attributevalue.MarshalMap(&primaryKey)

	res, err := db.GetItem(
		ctx,
		&dynamodb.GetItemInput{
			TableName: &cfg.Cfg().TableName,
			Key:       marshaledKey,
		},
	)
	if err != nil {
		return err
	}

	if res.Item != nil {
		return nil
	}

	resp, err := geo.GeocodeAddress(address)
	if err != nil {
		return err
	}

	practiceRecord := &PracticeRecord{
		PrimaryKey:  primaryKey,
		PracticeId:  id,
		Name:        name,
		FullAddress: address,
		Lattitude:   resp.Addresses[0].Latitude,
		Longitude:   resp.Addresses[0].Longitude,
		GeoHash:     "",
		Phone:       phone,
		Website:     website,
	}

	marshaledRecord, _ := attributevalue.MarshalMap(practiceRecord)

	_, err = db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &cfg.Cfg().TableName,
		Item:      marshaledRecord,
	})
	if err != nil {
		return err
	}

	return nil
}

func PutProvider(id, name, tags, practiceId string) error {
	return nil
}
