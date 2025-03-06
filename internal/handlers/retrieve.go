package handlers

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"lymphly/internal/data"
	"lymphly/internal/geo"
	"lymphly/internal/log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

var (
	enumerateAllPracticesFunc = data.EnumerateAllPractices
)

// Routes related to retrieval of data
func RetrieveRoutes(r chi.Router) {
	r.Get("/practices/all", GetAllPractices)
	r.Get("/practices/locate/proximity", ProximitySearch)
	r.Get("/practices/locate/state/{stateCode}", LocatePracticeByState)
	r.Get("/practice/{practiceId}", GetPractice)
	r.Get("/practice/{practiceId}/providers", GetPracticeByProviders)
	r.Get("/provider/{providerId}", GetProvider)
	r.Get("/provider/{providerId}/practice", GetPracticeByProvider)
}

// LimitedPracticeItem is a truncated amount of data about a practice
type LimitedPracticeItem struct {
	PracticeId string  `json:"practiceId,"`
	Name       string  `json:"name"`
	Lattitude  float64 `json:"lattitude"`
	Longitude  float64 `json:"longitude"`
}

// LimistedPracticeList is a list of limited practices
type LimitedPracticeList struct {
	Practices []LimitedPracticeItem `json:"practices"`
}

// Retrieves all Practices from the Database
func GetAllPractices(w http.ResponseWriter, r *http.Request) {
	d, err := enumerateAllPracticesFunc(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := &LimitedPracticeList{
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

// Retrieves a Single Practice from the database
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

// Response Body for all Providers for a Given Practice
type ProvidersByPracticeResponse struct {
	PracticeId string          `json:"practiceId"`
	Providers  []data.Provider `json:"providers"`
}

// Retrieves all Providers for a given practice Id
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

// Retrieves a given provider from the database
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

// Retrieves the Practice from a given provider id
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

// Given a State, lists all practices within
func LocatePracticeByState(w http.ResponseWriter, r *http.Request) {
	stateCode := chi.URLParam(r, "stateCode")
	if stateCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	stateCode = strings.ToUpper(stateCode)
	practices, err := data.EnumeratePracticesByState(r.Context(), stateCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := &LimitedPracticeList{
		Practices: make([]LimitedPracticeItem, len(practices)),
	}

	for idx, p := range practices {
		resp.Practices[idx] = LimitedPracticeItem{
			PracticeId: p.PracticeId,
			Name:       p.Name,
			Lattitude:  p.Lattitude,
			Longitude:  p.Longitude,
		}
	}

	outBytes, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(outBytes)
}

// Results of a Proximity Search
type ProximityResponse struct {
	*LimitedPracticeList
	OriginLatitude  float64 `json:"originLatitude,omitempty"`
	OriginLongitude float64 `json:"originLongitude,omitempty"`
}

// Searches in the proximity of a location for practices
func ProximitySearch(w http.ResponseWriter, r *http.Request) {

	/// Query Parameters
	// lat, long
	// addr

	latStr := r.URL.Query().Get("lat")
	longStr := r.URL.Query().Get("long")
	addrB64 := r.URL.Query().Get("addr")
	radius := r.URL.Query().Get("radius")
	radiusMi, err := strconv.Atoi(radius)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if radiusMi < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var lat float64
	var long float64

	// Process Lat/Long First
	if latStr != "" && longStr != "" {
		lat, err = strconv.ParseFloat(latStr, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		long, err = strconv.ParseFloat(longStr, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else if addrB64 != "" {
		addrBytes, err := base64.URLEncoding.DecodeString(addrB64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		addr := string(addrBytes)

		addrGeocode, err := geo.GeocodeAddress(addr)
		if errors.Is(err, geo.ErrBadAddress) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		lat = addrGeocode.Addresses[0].Latitude
		long = addrGeocode.Addresses[0].Longitude
	}

	practices, err := data.GetPracticesByProximity(r.Context(), lat, long, radiusMi)
	if err != nil {
		log.Error("failed to query practices by proximity", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := &LimitedPracticeList{
		Practices: make([]LimitedPracticeItem, len(practices)),
	}

	for idx, p := range practices {
		resp.Practices[idx] = LimitedPracticeItem{
			PracticeId: p.PracticeId,
			Name:       p.Name,
			Lattitude:  p.Lattitude,
			Longitude:  p.Longitude,
		}
	}

	outBytes, _ := json.Marshal(&ProximityResponse{
		OriginLatitude:      lat,
		OriginLongitude:     long,
		LimitedPracticeList: resp,
	})
	w.Header().Set("Content-Type", "application/json")
	w.Write(outBytes)
}
