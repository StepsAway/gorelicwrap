package gorelicwrap

import (
	"github.com/dghubble/sling"
	"net/http"
	"os"
)

// NewRelic defines the fields for storing NewRelic info necessary for making API requests
type NewRelic struct {
	client  *NRClient
	appID   string
	appName string
}

// NRClient allows for reuse of the NewRelic client
type NRClient struct {
	sling *sling.Sling
}

// NewNRClient returns an initialized NRClient
func NewNRClient(client *http.Client) *NRClient {
	baseURL := "https://api.newrelic.com/v2/"
	apiKey := os.Getenv("NEWRELIC_KEY")
	return &NRClient{
		sling: sling.New().Client(client).Base(baseURL).Set("X-Api-Key", apiKey),
	}
}
