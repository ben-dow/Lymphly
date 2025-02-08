package data

import (
	"context"
	"lymphly/internal/cfg"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func GetAllPractices(ctx context.Context) ([]Practice, error) {
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
