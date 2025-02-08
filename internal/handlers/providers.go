package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"lymphly/internal/data"
	"net/http"

	"golang.org/x/crypto/sha3"
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

	// Hasher for Determing Ids
	h := sha3.New512()

	// Determine Pratice Id
	practice := fmt.Sprintf("%s-%s", requestBody.Practice, requestBody.FullAddress)
	h.Write([]byte(practice))
	practiceId := base64.URLEncoding.EncodeToString(h.Sum(nil))

	// Reset Hasher
	h.Reset()

	// Determine Provider Id
	//provider := fmt.Sprintf("%s-%s", requestBody.Name, practiceId)
	//h.Write([]byte(provider))
	//providerId := base64.URLEncoding.EncodeToString(h.Sum(nil))

	// Save Practice
	err = data.PutPractice(r.Context(), practiceId, requestBody.Practice, requestBody.FullAddress, requestBody.Phone, requestBody.Website, requestBody.PracticeTags)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Save Provider
	//data.PutProvider(providerId, requestBody.Name, requestBody.ProviderTags, practiceId)

	w.WriteHeader(http.StatusOK)
}
