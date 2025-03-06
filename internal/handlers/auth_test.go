package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"lymphly/internal/auth"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	router := chi.NewRouter()
	AuthRoutes(router)

	t.Run("NoCredentials", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/login", nil)
		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("NoUsername", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/login", nil)
		r.Header.Set("x-username", "testUser")
		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("NoPassword", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/login", nil)
		r.Header.Set("x-password", "testPassword")
		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("FailGetAccessTokenPassReset", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/login", nil)
		r.Header.Set("x-username", "testUser")
		r.Header.Set("x-password", "testPassword")

		getAccessTokenFunc = func(ctx context.Context, username, password string) (string, error) {
			return "", auth.ErrPasswordReset
		}

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("FailGetAccessTokenCtxCanceled", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/login", nil)
		r.Header.Set("x-username", "testUser")
		r.Header.Set("x-password", "testPassword")

		getAccessTokenFunc = func(ctx context.Context, username, password string) (string, error) {
			return "", context.Canceled
		}

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("FailureUnkown", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/login", nil)
		r.Header.Set("x-username", "testUser")
		r.Header.Set("x-password", "testPassword")

		getAccessTokenFunc = func(ctx context.Context, username, password string) (string, error) {
			return "", errors.New("error")
		}

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/login", nil)
		r.Header.Set("x-username", "testUser")
		r.Header.Set("x-password", "testPassword")

		getAccessTokenFunc = func(ctx context.Context, username, password string) (string, error) {
			return "returnedAccessToken", nil
		}

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, "no-cache", w.Result().Header.Get("Cache-Control"))
		assert.Equal(t, "application/json", w.Result().Header.Get("Content-Type"))

		resp, _ := io.ReadAll(w.Body)
		assert.NotEmpty(t, resp)
		out := map[string]string{}
		err := json.Unmarshal(resp, &out)
		assert.NoError(t, err)
		assert.Equal(t, out["token"], "returnedAccessToken")
	})

}
