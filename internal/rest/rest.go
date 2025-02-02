package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRestServer() *chi.Mux {
	r := chi.NewRouter()
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(r.URL.RawPath))
	})
	r.Route("/api/v1/providersearch", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	})

	return r
}
