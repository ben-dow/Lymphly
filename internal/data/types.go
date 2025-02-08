package data

type PrimaryKey struct {
	PartitionKey string `dynamodbav:"pk"`
	SortKey      string `dynamodbav:"sk"`
}

const (
	ProvidersPk string = "provider"
)

type Provider struct {
	ProviderId string
	Name       string
	PracticeId string
	Tags       []string
}

type ProviderRecord struct {
	PrimaryKey
	Provider
}

const (
	PracticesPk string = "practice"
)

type Practice struct {
	PracticeId  string
	Name        string
	FullAddress string
	Lattitude   float64
	Longitude   float64
	GeoHash     string
	Phone       string
	Website     string
	State       string
	StateCode   string
	Country     string
	CountryCode string
	Tags        []string
}

type PracticeRecord struct {
	PrimaryKey
	Practice
}

const (
	PracticeGeoHashPkPrefix string = "practicehash#"
	ProviderGeoHashPkPrefix string = "providerhash#"
)

type PracticeGeoHashRecord struct {
	PrimaryKey
	Practice
}

type ProviderGeoHashRecord struct {
	PrimaryKey
	Provider
}
