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

func TestPutNewProvider(t *testing.T) {
	router := chi.NewRouter()
	ModificationRoutes(router)

	t.Run("InvalidBody", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("PUT", "/provider", strings.NewReader("body"))
		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("FailPutPractice", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("PUT", "/provider", strings.NewReader("{}"))

		putPracticeFunc = func(ctx context.Context, name, address, phone, website, tags string) (*data.Practice, error) {
			return nil, errors.New("err")
		}

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("FailPutProvider", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("PUT", "/provider", strings.NewReader("{}"))

		putPracticeFunc = func(ctx context.Context, name, address, phone, website, tags string) (*data.Practice, error) {
			return &data.Practice{}, nil
		}

		putProviderFunc = func(ctx context.Context, name, tags string, practice *data.Practice) (*data.Provider, error) {
			return nil, errors.New("err")
		}

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("FailPutProvider", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("PUT", "/provider", strings.NewReader("{}"))

		putPracticeFunc = func(ctx context.Context, name, address, phone, website, tags string) (*data.Practice, error) {
			return &data.Practice{}, nil
		}

		putProviderFunc = func(ctx context.Context, name, tags string, practice *data.Practice) (*data.Provider, error) {
			return &data.Provider{}, nil
		}

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
}
