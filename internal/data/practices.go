package data

import (
	"context"
	"errors"
	"lymphly/internal/cfg"
	"lymphly/internal/geo"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/mmcloughlin/geohash"
	"golang.org/x/sync/errgroup"
)

func PutPractice(ctx context.Context, name, address, phone, website, tags string) (*Practice, error) {
	// Geocode the Address
	geoCodeResp, err := geo.GeocodeAddress(address)
	if err != nil {
		return nil, err
	}
	addr := geoCodeResp.Addresses[0]

	// Determine Practice Id from Fully Formatted Addres from Geocode
	practiceId := DerivePracticeId(name, addr.FormattedAddress)

	// Check if Practice Exists
	practice, err := GetPractice(ctx, practiceId)
	if err == nil {
		return practice, nil
	} else if !errors.Is(err, ErrPracticeNotFound) {
		return nil, err
	}

	// If Practice does not exist, add it to the database

	// Practice Definition
	practice = &Practice{
		PracticeId:  practiceId,
		Name:        name,
		State:       addr.State,
		StateCode:   addr.StateCode,
		Country:     addr.Country,
		CountryCode: addr.CountryCode,
		FullAddress: addr.FormattedAddress,
		Lattitude:   addr.Latitude,
		Longitude:   addr.Longitude,
		GeoHash:     geohash.EncodeWithPrecision(addr.Latitude, addr.Longitude, 4),
		Phone:       phone,
		Website:     website,
	}

	execgroup := new(errgroup.Group)

	// Save Practice Record
	execgroup.Go(func() error {
		practiceRecord := &PracticeRecord{
			PrimaryKey: NewPracticePrimaryKey(practiceId),
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

	// Save Practice Geohash Record
	execgroup.Go(func() error {
		practiceRecord := &PracticeGeoHashRecord{
			PrimaryKey: NewPracticeGeoHashPrimaryKey(practice.GeoHash, practiceId),
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

	// Wait for saving to database to complete
	err = execgroup.Wait()
	if err != nil {
		return nil, err
	}

	return practice, nil
}

func EnumerateAllPractices(ctx context.Context) ([]Practice, error) {
	keyCond := expression.Key("pk").Equal(expression.Value("practices"))
	expr, _ := expression.NewBuilder().WithKeyCondition(keyCond).Build()

	paginator := dynamodb.NewQueryPaginator(db, &dynamodb.QueryInput{
		TableName:                 aws.String(cfg.Cfg().TableName),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	results := []map[string]types.AttributeValue{}
	for paginator.HasMorePages() {
		res, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		results = append(results, res.Items...)
	}

	out := make([]Practice, len(results))
	for idx, r := range results {
		res := Practice{}
		err := attributevalue.UnmarshalMap(r, &res)
		if err != nil {
			return nil, err
		}
		out[idx] = res
	}

	return out, nil
}

var ErrPracticeNotFound = errors.New("practice not found")

func GetPractice(ctx context.Context, practiceId string) (*Practice, error) {
	marshaledKey, _ := attributevalue.MarshalMap(NewPracticePrimaryKey(practiceId))

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

	if res.Item == nil {
		return nil, ErrPracticeNotFound
	}

	out := &PracticeRecord{}
	err = attributevalue.UnmarshalMap(res.Item, out)
	if err != nil {
		return nil, err
	}
	return out.Practice, nil
}

func EnumeratePracticesByState(ctx context.Context, stateCode string) ([]Practice, error) {
	keyCond := expression.Key("pk").Equal(expression.Value(PracticesPk))
	cond := expression.Name("StateCode").Equal(expression.Value(stateCode))
	expr, _ := expression.NewBuilder().WithKeyCondition(keyCond).WithFilter(cond).Build()

	paginator := dynamodb.NewQueryPaginator(db, &dynamodb.QueryInput{
		TableName:                 aws.String(cfg.Cfg().TableName),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	results := []map[string]types.AttributeValue{}
	for paginator.HasMorePages() {
		res, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		results = append(results, res.Items...)
	}

	out := make([]Practice, len(results))
	for idx, r := range results {
		res := Practice{}
		err := attributevalue.UnmarshalMap(r, &res)
		if err != nil {
			return nil, err
		}
		out[idx] = res
	}

	return out, nil
}

func EnumeratePracticesByGeoHash(ctx context.Context, hashes []string) ([]Practice, error) {

	resMu := sync.Mutex{}
	outResults := []Practice{}
	execGroup := new(errgroup.Group)

	for i := 0; i < len(hashes); i += 100 {
		execGroup.Go(func() error {
			end := i + 100
			if end > len(hashes) {
				end = len(hashes)
			}
			hashChunk := hashes[i:end]
			keyCond := expression.Key("pk").Equal(expression.Value(PracticesPk))
			hashValues := make([]expression.OperandBuilder, len(hashChunk))
			for idx, h := range hashChunk {
				hashValues[idx] = expression.Value(h)
			}
			cond := expression.Name("GeoHash").In(hashValues[0], hashValues[1:]...)

			expr, _ := expression.NewBuilder().WithKeyCondition(keyCond).WithFilter(cond).Build()

			paginator := dynamodb.NewQueryPaginator(db, &dynamodb.QueryInput{
				TableName:                 aws.String(cfg.Cfg().TableName),
				KeyConditionExpression:    expr.KeyCondition(),
				FilterExpression:          expr.Filter(),
				ProjectionExpression:      expr.Projection(),
				ExpressionAttributeNames:  expr.Names(),
				ExpressionAttributeValues: expr.Values(),
			})

			results := []map[string]types.AttributeValue{}
			for paginator.HasMorePages() {
				res, err := paginator.NextPage(ctx)
				if err != nil {
					return err
				}
				results = append(results, res.Items...)
			}

			for _, r := range results {
				res := Practice{}
				err := attributevalue.UnmarshalMap(r, &res)
				if err != nil {
					return err
				}
				resMu.Lock()
				outResults = append(outResults, res)
				resMu.Unlock()
			}

			return nil
		})
	}

	err := execGroup.Wait()
	if err != nil {
		return nil, err
	}

	return outResults, nil
}

func GetPracticesByProximity(ctx context.Context, lat, long float64, radius int) ([]Practice, error) {
	originHash := geohash.EncodeWithPrecision(lat, long, 4)

	depth := 0
	if radius <= 20 {
		depth = 1
	} else if radius <= 50 {
		depth = 2
	} else if radius <= 100 {
		depth = 5
	} else {
		return nil, errors.New("cant do more than 100 miles")
	}

	hashes := geo.Neighbors(originHash, depth)

	practices, err := EnumeratePracticesByGeoHash(ctx, hashes)
	if err != nil {
		return nil, err
	}

	mu := sync.Mutex{}
	out := []Practice{}
	wg := sync.WaitGroup{}
	for _, p := range practices {
		wg.Add(1)
		go func() {
			if _, in := geo.InRadius(lat, long, p.Lattitude, p.Longitude, radius); in {
				mu.Lock()
				out = append(out, p)
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return out, nil
}
