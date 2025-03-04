package main

import (
	"lymphly/internal/cfg"
	"lymphly/internal/handlers"

	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	chi "github.com/go-chi/chi/v5"
)

// This is the entry function into the lambda handler for the private endpoint of the Lymphly API
// The Lymphly API utilizes Chi as a router and the lambda-go-api-proxy to convert from Lambda requests
// fed by an API Gateway to HTTP requests that are able to be handled by the Chi router and Std Library HTTP Routes
//
// All routes in this API are expected to be cached
func main() {
	r := chi.NewRouter()
	r.Route(cfg.Cfg().BasePath, func(r chi.Router) {
		handlers.GeneralRoutes(r)
		handlers.RetrieveRoutes(r)
		handlers.AuthRoutes(r)
	})

	adapter := chiadapter.New(r)
	lambda.Start(adapter.ProxyWithContext)
}
