package main

import (
	"lymphly/internal/cfg"
	"lymphly/internal/handlers"

	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	chi "github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Route(cfg.Cfg().BasePath, func(r chi.Router) {
		handlers.GeneralRoutes(r)
		r.Put("/provider", handlers.PutNewProvider)
	})

	adapter := chiadapter.New(r)
	lambda.Start(adapter.ProxyWithContext)
}
