package handlers

import (
	"context"
	"errors"
	"fmt"
	"lymphly/internal/data"
	"lymphly/internal/geo"
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

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practices/all", strings.NewReader("body"))
		enumerateAllPracticesFunc = func(ctx context.Context) ([]data.Practice, error) {
			return []data.Practice{{}}, nil
		}
		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
}

func TestGetPractice(t *testing.T) {
	router := chi.NewRouter()
	RetrieveRoutes(router)

	t.Run("MissingPracticeId", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practice/practiceId", strings.NewReader("body"))

		GetPractice(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("GetPracticeFailure", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practice/practiceId", strings.NewReader("body"))
		getPracticeFunc = func(ctx context.Context, practiceId string) (*data.Practice, error) {
			return nil, errors.New("err")
		}

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practice/practiceId", strings.NewReader("body"))
		getPracticeFunc = func(ctx context.Context, practiceId string) (*data.Practice, error) {
			return &data.Practice{}, nil
		}

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
}

func TestGetPracticeProviders(t *testing.T) {
	router := chi.NewRouter()
	RetrieveRoutes(router)

	t.Run("MissingPracticeId", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practice/practiceId/providers", strings.NewReader("body"))

		GetPracticeProviders(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("GetProvidersFailure", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practice/practiceId/providers", strings.NewReader("body"))

		getProvidersByPracticeIdFunc = func(ctx context.Context, practiceId string) ([]data.Provider, error) {
			return nil, errors.New("err")
		}

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practice/practiceId/providers", strings.NewReader("body"))
		getProvidersByPracticeIdFunc = func(ctx context.Context, practiceId string) ([]data.Provider, error) {
			return []data.Provider{{}}, nil
		}

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
}

func TestGetProvider(t *testing.T) {
	router := chi.NewRouter()
	RetrieveRoutes(router)

	t.Run("MissingProviderId", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/provider/providerId", nil)

		GetProvider(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("FailGetProvider", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/provider/providerId", nil)

		getProviderFunc = func(ctx context.Context, providerId string) (*data.Provider, error) {
			return nil, fmt.Errorf("err")
		}

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/provider/providerId", nil)
		getProviderFunc = func(ctx context.Context, providerId string) (*data.Provider, error) {
			return &data.Provider{}, nil
		}
		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
}

func TestGetPracticeByProvider(t *testing.T) {
	router := chi.NewRouter()
	RetrieveRoutes(router)

	t.Run("MissingProviderId", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/provider/providerId/practice", nil)

		GetPracticeByProvider(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("FailGetProvider", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/provider/providerId/practice", nil)

		getProviderFunc = func(ctx context.Context, providerId string) (*data.Provider, error) {
			return nil, fmt.Errorf("err")
		}

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("FailGetPractice", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/provider/providerId/practice", nil)

		getProviderFunc = func(ctx context.Context, providerId string) (*data.Provider, error) {
			return &data.Provider{}, nil
		}

		getPracticeFunc = func(ctx context.Context, practiceId string) (*data.Practice, error) {
			return nil, fmt.Errorf("err")
		}

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/provider/providerId/practice", nil)
		getProviderFunc = func(ctx context.Context, providerId string) (*data.Provider, error) {
			return &data.Provider{}, nil
		}
		getPracticeFunc = func(ctx context.Context, practiceId string) (*data.Practice, error) {
			return &data.Practice{}, nil
		}

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
}

func TestEnumeratePracticeByState(t *testing.T) {
	router := chi.NewRouter()
	RetrieveRoutes(router)

	t.Run("MissingStateCode", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practices/locate/state/me", nil)

		LocatePracticeByState(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("EnumerateFailure", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practices/locate/state/me", nil)
		enumeratePracticesByStateFunc = func(ctx context.Context, stateCode string) ([]data.Practice, error) {
			return nil, fmt.Errorf("err")
		}
		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practices/locate/state/me", nil)
		enumeratePracticesByStateFunc = func(ctx context.Context, stateCode string) ([]data.Practice, error) {
			return []data.Practice{{}}, nil
		}
		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
}

func TestProximityBadRadius(t *testing.T) {
	router := chi.NewRouter()
	RetrieveRoutes(router)

	t.Run("BadRadius", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practices/locate/proximity?radius=notint", nil)

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("SmallRadius", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practices/locate/proximity?radius=-10", nil)

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})
}

func TestProximityLatLong(t *testing.T) {
	router := chi.NewRouter()
	RetrieveRoutes(router)

	t.Run("BadLat", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practices/locate/proximity?radius=10&lat=notfloat&long=notfloat", nil)

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("BadLong", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practices/locate/proximity?radius=10&lat=10&long=notfloat", nil)

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("FailQuery", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practices/locate/proximity?radius=10&lat=10&long=10", nil)

		getPracticesByProximityFunc = func(ctx context.Context, lat, long float64, radius int) ([]data.Practice, error) {
			return nil, fmt.Errorf("err")
		}

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practices/locate/proximity?radius=10&lat=10&long=10", nil)

		getPracticesByProximityFunc = func(ctx context.Context, lat, long float64, radius int) ([]data.Practice, error) {
			return []data.Practice{{}}, nil
		}

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
}

func TestProxmityAddr(t *testing.T) {
	router := chi.NewRouter()
	RetrieveRoutes(router)

	t.Run("BadAddr", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practices/locate/proximity?radius=10&addr=*", nil)

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("FailQuery", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practices/locate/proximity?radius=10&addr=addr", nil)

		geocodeAddrFunc = func(addr string) (*geo.GeocodeResponse, error) {
			return nil, geo.ErrBadAddress
		}

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("FailQuery", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practices/locate/proximity?radius=10&addr=addr", nil)

		geocodeAddrFunc = func(addr string) (*geo.GeocodeResponse, error) {
			return nil, fmt.Errorf("err")
		}

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/practices/locate/proximity?radius=10&addr=addr", nil)

		geocodeAddrFunc = func(addr string) (*geo.GeocodeResponse, error) {
			return &geo.GeocodeResponse{
				Addresses: []geo.AddressResponse{{}},
			}, nil
		}

		getPracticesByProximityFunc = func(ctx context.Context, lat, long float64, radius int) ([]data.Practice, error) {
			return []data.Practice{{}}, nil
		}

		router.ServeHTTP(w, r)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
}
