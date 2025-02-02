package main

import (
	"lymphly/internal/routing"

	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	chi "github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	routing.GeneralRoutes(r)

	adapter := chiadapter.New(r)
	lambda.Start(adapter.ProxyWithContext)
}
