package main

import (
	"lymphly/internal/cfg"
	"lymphly/internal/handlers"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	chi "github.com/go-chi/chi/v5"
)

// This is the entry function into the lambda handler for the private endpoint of the Lymphly API
// The Lymphly API utilizes Chi as a router and the lambda-go-api-proxy to convert from Lambda requests
// fed by an API Gateway to HTTP requests that are able to be handled by the Chi router and Std Library HTTP Routes
//
// All requests to the endpoints defined in this file are expected to have been authenticated by the API Gateway authorizer
// All responses will have the Cache-Control: no-cache header to prevent caching and allow for live editing of data
func main() {
	r := chi.NewRouter()
	r.Route(cfg.Cfg().BasePath, func(r chi.Router) {
		r.Use(func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Cache-Control", "no-cache")
				h.ServeHTTP(w, r)
			})
		})
		handlers.GeneralRoutes(r)
		handlers.RetrieveRoutes(r)
		handlers.ModificationRoutes(r)
	})

	adapter := chiadapter.New(r)
	lambda.Start(adapter.ProxyWithContext)
}
