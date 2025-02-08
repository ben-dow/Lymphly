package data

type PrimaryKey struct {
	PartitionKey string `dynamodbav:"pk"`
	SortKey      string `dynamodbav:"sk"`
}

const (
	ProviderPk string = "providers"
)

type Provider struct {
	ProviderId string   `json:"providerId,omitempty"`
	Name       string   `json:"name,omitempty"`
	PracticeId string   `json:"practiceId,omitempty"`
	Tags       []string `json:"tags,omitempty"`
}

type ProviderRecord struct {
	PrimaryKey
	Provider
}

const (
	PracticesPk             string = "practices"
	PracticeGeoHashPkPrefix string = "practicehash#"
)

type Practice struct {
	PracticeId  string   `json:"practiceId,omitempty"`
	Name        string   `json:"name,omitempty"`
	FullAddress string   `json:"fullAddress,omitempty"`
	Lattitude   float64  `json:"lattitude,omitempty"`
	Longitude   float64  `json:"longitude,omitempty"`
	GeoHash     string   `json:"geoHash,omitempty"`
	Phone       string   `json:"phone,omitempty"`
	Website     string   `json:"website,omitempty"`
	State       string   `json:"state,omitempty"`
	StateCode   string   `json:"stateCode,omitempty"`
	Country     string   `json:"country,omitempty"`
	CountryCode string   `json:"countryCode,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

type PracticeRecord struct {
	PrimaryKey
	Practice
}

type PracticeGeoHashRecord struct {
	PrimaryKey
	Practice
}
