package handlers

import (
	"encoding/json"
	"lymphly/internal/data"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RetrieveRoutes(r chi.Router) {
	r.Get("/practices/all", AllPractices)
	r.Get("/practices/locate/proximity", LocatePractice)
	r.Get("/practices/locate/state/{stateCode}", LocatePracticeByState)
	r.Get("/practice/{practiceId}", GetPractice)
	r.Get("/practice/{practiceId}/providers", GetPracticeByProviders)
	r.Get("/provider/{providerId}", GetProvider)
	r.Get("/provider/{providerId}/practice", GetPracticeByProvider)
}

type LimitedPracticeItem struct {
	PracticeId string  `json:"practiceId,"`
	Name       string  `json:"name"`
	Lattitude  float64 `json:"lattitude"`
	Longitude  float64 `json:"longitude"`
}

type AllPracticesResponse struct {
	Practices []LimitedPracticeItem `json:"practices"`
}

func AllPractices(w http.ResponseWriter, r *http.Request) {
	d, err := data.EnumerateAllPractices(r.Context())
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

func GetPractice(w http.ResponseWriter, r *http.Request) {
	practiceId := chi.URLParam(r, "practiceId")
	if practiceId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	practice, err := data.GetPractice(r.Context(), practiceId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	practiceBytes, _ := json.Marshal(practice)
	w.Header().Set("Content-Type", "application/json")
	w.Write(practiceBytes)
}

type ProvidersByPracticeResponse struct {
	PracticeId string          `json:"practiceId"`
	Providers  []data.Provider `json:"providers"`
}

func GetPracticeByProviders(w http.ResponseWriter, r *http.Request) {
	practiceId := chi.URLParam(r, "practiceId")
	if practiceId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	providers, err := data.GetProvidersByPracticeId(r.Context(), practiceId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := &ProvidersByPracticeResponse{
		PracticeId: practiceId,
		Providers:  providers,
	}

	respBytes, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respBytes)
}

func GetProvider(w http.ResponseWriter, r *http.Request) {
	providerId := chi.URLParam(r, "providerId")
	if providerId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	provider, err := data.GetProvider(r.Context(), providerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBytes, _ := json.Marshal(provider)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBytes)
}

func GetPracticeByProvider(w http.ResponseWriter, r *http.Request) {
	providerId := chi.URLParam(r, "providerId")
	if providerId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	provider, err := data.GetProvider(r.Context(), providerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	practice, err := data.GetPractice(r.Context(), provider.PracticeId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	practiceBytes, _ := json.Marshal(practice)
	w.Header().Set("Content-Type", "application/json")
	w.Write(practiceBytes)
}

type EnumeratedPractices struct {
	Practices []data.Practice `json:"practices"`
}

func LocatePracticeByState(w http.ResponseWriter, r *http.Request) {
	stateCode := chi.URLParam(r, "stateCode")
	if stateCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	practices, err := data.EnumeratePracticesByState(r.Context(), stateCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := &EnumeratedPractices{
		Practices: practices,
	}

	outBytes, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(outBytes)
}

func LocatePractice(w http.ResponseWriter, r *http.Request) {
	/// Query Parameters
	// lat, long
	// addr

}
