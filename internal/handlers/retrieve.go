package handlers

import (
	"encoding/json"
	"lymphly/internal/data"
	"net/http"
)

type ListPracticesItem struct {
	PracticeId  string   `json:"practice_id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Lattitude   float64  `json:"lattitude,omitempty"`
	Longitude   float64  `json:"longitude,omitempty"`
	Phone       string   `json:"phone,omitempty"`
	Website     string   `json:"website,omitempty"`
	State       string   `json:"state,omitempty"`
	StateCode   string   `json:"stateCode,omitempty"`
	CountryCode string   `json:"countryCode,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

type ListPracticesResponse struct {
	Practices []ListPracticesItem `json:"practices,omitempty"`
}

func ListPractices(w http.ResponseWriter, r *http.Request) {
	d, err := data.GetAllPractices(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := &ListPracticesResponse{
		Practices: make([]ListPracticesItem, len(d)),
	}
	for idx, p := range d {
		response.Practices[idx] = ListPracticesItem{
			PracticeId:  p.PracticeId,
			Name:        p.Name,
			Lattitude:   p.Lattitude,
			Longitude:   p.Longitude,
			Phone:       p.Phone,
			Website:     p.Website,
			State:       p.State,
			StateCode:   p.StateCode,
			CountryCode: p.CountryCode,
			Tags:        p.Tags,
		}
	}

	outBytes, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(outBytes)
}
