package geo

import (
	"encoding/json"
	"errors"
	"io"
	"lymphly/internal/cfg"
	"net/http"
	"net/url"
)

var httpClient = &http.Client{}

type MetaResponse struct {
	Code int `json:"code"`
}

type AddressResponse struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Geometry  struct {
		Type        string    `json:"type"`
		Coordinates []float32 `json:"coordinates"`
	} `json:"geometry"`
	Country          string `json:"country"`
	CountryCode      string `json:"countryCode"`
	CountryFlag      string `json:"countryFlag"`
	County           string `json:"county"`
	Confidence       string `json:"confidence"`
	Borough          string `json:"borough"`
	City             string `json:"city"`
	Number           string `json:"number"`
	Neighborhood     string `json:"neighborhood"`
	PostalCode       string `json:"postalCode"`
	StateCode        string `json:"stateCode"`
	State            string `json:"state"`
	Street           string `json:"street"`
	Layer            string `json:"layer"`
	FormattedAddress string `json:"formattedAddress"`
	AddressLabel     string `json:"addressLabel"`
	TimeZone         struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Code        string `json:"code"`
		CurrentTime string `json:"currentTime"`
		DstOffset   int    `json:"dstOffset"`
		UtcOffset   int    `json:"utcOffset"`
	} `json:"timeZone"`
}

type GeocodeResponse struct {
	Meta      MetaResponse      `json:"meta,omitempty"`
	Addresses []AddressResponse `json:"addresses,omitempty"`
}

func GeocodeAddress(addr string) (*GeocodeResponse, error) {

	req, err := http.NewRequest("GET", "https://api.radar.io/v1/geocode/forward?query="+url.QueryEscape(addr), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", cfg.Cfg().RadarPrivateKey)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("request was not successful")
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	out := &GeocodeResponse{}
	err = json.Unmarshal(respBytes, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
