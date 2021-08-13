package testutil

import (
	"net/http"

	"github.com/SeNicko/lvlup"
)

type roundTripFunc func(req *http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

// NewTestClient creates a new client with mocked http client.
func NewTestLvlClient(apiKey string, handler roundTripFunc) *lvlup.LvlClient {
	httpClient := &http.Client{
		Transport: roundTripFunc(handler),
	}
	client := lvlup.NewLvlClient(apiKey, httpClient)
	return client
}
