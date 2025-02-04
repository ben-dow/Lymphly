package handlers

type Provider struct {
	Id           string
	PracticeId   string
	Name         string
	ProviderType string
}

type Practice struct {
	Id          string
	Name        string
	FullAddress string
	Geocode     string
	Phone       string
	Website     string
}

