package handlers

import (
	"context"
	"errors"
	"lymphly/internal/data"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestGetAllPractices(t *testing.T) {
	router := chi.NewRouter()
	RetrieveRoutes(router)

	t.Run("EnumerateFailure", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practices/all", strings.NewReader("body"))
		enumerateAllPracticesFunc = func(ctx context.Context) ([]data.Practice, error) {
			return nil, errors.New("error")
		}
		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}
