package testutil

import (
	"net/http"
)

// Error returns http.HandlerFunc which sets status code to provided value.
func HttpError(status int) roundTripFunc {
	return func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: status,
		}, nil
	}
}
