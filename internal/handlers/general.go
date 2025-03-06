package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// GeneralRoutes are non-domain routes for information about the service
func GeneralRoutes(r chi.Router) {
	r.Get("/health", Health)
}

// Health returns a 200 if the service is okay. 
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")
	w.WriteHeader(http.StatusOK)
}
