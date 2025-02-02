package routing

import (
	"encoding/json"
	"lymphly/internal/cfg"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GeneralRoutes(r chi.Router) {
	r.Get("/health", Health)
	r.Get("/manifest", Manifest)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

type ManifestResponse struct {
	Version string
}

func Manifest(w http.ResponseWriter, r *http.Request) {
	resp := &ManifestResponse{
		Version: cfg.Cfg().Version,
	}
	b, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
