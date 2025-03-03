package main

import (
	"lymphly/internal/cfg"
	"lymphly/internal/handlers"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	chi "github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Route(cfg.Cfg().BasePath, func(r chi.Router) {
		r.Use(func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Cache-Control", "no-cache")
			})
		})
		handlers.GeneralRoutes(r)
		r.Put("/provider", handlers.PutNewProvider)
	})

	adapter := chiadapter.New(r)
	lambda.Start(adapter.ProxyWithContext)
}
