package data

type PrimaryKey struct {
	PartitionKey string `dynamodbav:"pk"`
	SortKey      string `dynamodbav:"sk"`
}

const (
	ProvidersPk string = "provider"
)

type ProviderRecord struct {
	PrimaryKey
	ProviderId string
	Name       string
	Type       string
	PracticeId string
}

const (
	PracticesPk string = "practice"
)

type PracticeRecord struct {
	PrimaryKey
	PracticeId  string
	Name        string
	FullAddress string
	Lattitude   float32
	Longitude   float32
	GeoHash     string
	Phone       string
	Website     string
}

const (
	GeoHashPkPrefix    string = "geohash#"
	GeoHashSkPractices string = "practices"
)

type PracticeGeoHash struct {
	PrimaryKey
	Practices []string
}
