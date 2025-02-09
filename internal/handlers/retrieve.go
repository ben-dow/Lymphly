package handlers

import (
	"encoding/json"
	"lymphly/internal/data"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RetrieveRoutes(r chi.Router) {
	r.Get("/practices/all", Health)
}

type LimitedPracticeItem struct {
	PracticeId string  `json:"practiceId,"`
	Name       string  `json:"name"`
	Lattitude  float64 `json:"lattitude"`
	Longitude  float64 `json:"longitude"`
}

type AllPracticesResponse struct {
	Practices []LimitedPracticeItem `json:"practices,omitempty"`
}

func AllPractices(w http.ResponseWriter, r *http.Request) {
	d, err := data.QueryPractices(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := &AllPracticesResponse{
		Practices: make([]LimitedPracticeItem, len(d)),
	}

	for idx, p := range d {
		response.Practices[idx] = LimitedPracticeItem{
			PracticeId: p.PracticeId,
			Name:       p.Name,
			Lattitude:  p.Lattitude,
			Longitude:  p.Longitude,
		}
	}

	outBytes, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(outBytes)
}
