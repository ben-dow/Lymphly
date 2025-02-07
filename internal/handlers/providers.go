package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
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

	// Determine Pratice Id
	practice := fmt.Sprintf("%s-%s", requestBody.Practice, requestBody.FullAddress)
	h := sha3.New512()
	h.Write([]byte(practice))
	id := base64.URLEncoding.EncodeToString(h.Sum(nil))

	w.Write([]byte(id))

}
