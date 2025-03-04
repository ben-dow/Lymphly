package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"lymphly/internal/auth"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// AuthRoutes populates the provided router with routes/handlers related to authentication
func AuthRoutes(r chi.Router) {
	r.Get("/login", GetLogin)
}

// GetLogin handles GET requests to login and retrieve a JWT token
func GetLogin(w http.ResponseWriter, r *http.Request) {

	// Parse Credentials from Request
	username := r.Header.Get("x-username")
	password := r.Header.Get("x-password")
	if username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get Token
	accessToken, err := auth.GetAccessToken(r.Context(), username, password)
	if err != nil {
		if errors.Is(err, auth.ErrPasswordReset) || errors.Is(err, context.Canceled) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Prepare Response
	out := map[string]string{"token": accessToken}
	outBytes, _ := json.Marshal(out)

	// Write Response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(outBytes)
}
