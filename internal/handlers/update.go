package handlers

import (
	"encoding/json"
	"io"
	"lymphly/internal/data"
	"net/http"
)

type NewProviderRequest struct {
	Name         string `json:"name"`
	Practice     string `json:"practice"`
	FullAddress  string `json:"fullAddress"`
	Phone        string `json:"phone"`
	Website      string `json:"website"`
	ProviderTags string `json:"providerTags"`
	PracticeTags string `json:"practiceTags"`
}

func PutNewProvider(w http.ResponseWriter, r *http.Request) {
	// Ready Body
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Parse Request
	requestBody := &NewProviderRequest{}
	err = json.Unmarshal(b, requestBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Save Practice
	savedPractice, err := data.PutPractice(r.Context(), requestBody.Practice, requestBody.FullAddress, requestBody.Phone, requestBody.Website, requestBody.PracticeTags)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Save Provider
	_, err = data.PutProvider(r.Context(), requestBody.Name, requestBody.ProviderTags, savedPractice)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
