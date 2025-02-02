package main

import (
	"lymphly/internal/rest"

	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
)

func main() {
	r := rest.NewRestServer()
	adapter := chiadapter.New(r)
	lambda.Start(adapter.Proxy)
}
