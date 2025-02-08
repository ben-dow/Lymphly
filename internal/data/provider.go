package data

import (
	"context"
	"fmt"
	"lymphly/internal/cfg"
	"lymphly/internal/geo"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/mmcloughlin/geohash"
	"golang.org/x/sync/errgroup"
)

var db *dynamodb.Client

func init() {
	cfg, _ := config.LoadDefaultConfig(context.Background())
	db = dynamodb.NewFromConfig(cfg)
}

func PutPractice(ctx context.Context, id, name, address, phone, website, tags string) (*Practice, error) {

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
		return nil, err
	}

	if res.Item != nil {
		out := &PracticeRecord{}
		err := attributevalue.UnmarshalMap(res.Item, out)
		if err != nil {
			return nil, err
		}
		return &out.Practice, nil
	}

	resp, err := geo.GeocodeAddress(address)
	if err != nil {
		return nil, err
	}

	practice := Practice{
		PracticeId:  id,
		Name:        name,
		State:       resp.Addresses[0].State,
		StateCode:   resp.Addresses[0].StateCode,
		Country:     resp.Addresses[0].Country,
		CountryCode: resp.Addresses[0].CountryCode,
		FullAddress: resp.Addresses[0].FormattedAddress,
		Lattitude:   resp.Addresses[0].Latitude,
		Longitude:   resp.Addresses[0].Longitude,
		GeoHash:     geohash.EncodeWithPrecision(resp.Addresses[0].Latitude, resp.Addresses[0].Longitude, 4),
		Phone:       phone,
		Website:     website,
	}

	execgroup := new(errgroup.Group)

	execgroup.Go(func() error {
		practiceRecord := &PracticeRecord{
			PrimaryKey: primaryKey,
			Practice:   practice,
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
	})

	execgroup.Go(func() error {
		practiceRecord := &PracticeGeoHashRecord{
			PrimaryKey: PrimaryKey{
				PartitionKey: fmt.Sprintf("%s%s", PracticeGeoHashPkPrefix, practice.GeoHash),
				SortKey:      practice.PracticeId,
			},
			Lattitude:  practice.Lattitude,
			Longitude:  practice.Longitude,
			PracticeId: practice.PracticeId,
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
	})

	err = execgroup.Wait()
	if err != nil {
		return nil, err
	}

	return &practice, nil
}

func PutProvider(id, name, tags, practiceId string) error {
	return nil
}
