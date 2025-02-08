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
	Type       string
	PracticeId string
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
