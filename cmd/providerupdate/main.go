package main

import (
	"lymphly/internal/routes"

	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	chi "github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Route("/api/v1/providerupdate", func(r chi.Router) {
		r.Get("/health", routes.Health)
	})

	adapter := chiadapter.New(r)
	lambda.Start(adapter.ProxyWithContext)
}
