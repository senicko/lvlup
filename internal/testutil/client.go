package testutil

import (
	"net/http"

	"github.com/senicko/lvlup"
)

type RoundTripFunc func(req *http.Request) (*http.Response, error)

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

// NewTestLvlClient creates a new client with mocked http client.
func NewTestLvlClient(apiKey string, handler RoundTripFunc) *lvlup.LvlClient {
	httpClient := &http.Client{
		Transport: handler,
	}

	client := lvlup.NewLvlClient(apiKey, httpClient)
	return client
}

// HttpError returns http.HandlerFunc which sets status code to provided value.
func HttpError(status int) RoundTripFunc {
	return func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: status,
		}, nil
	}
}
