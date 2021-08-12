package testutil

import (
	"encoding/json"
	"net/http"
)

// SuccessJSON returns a http.HandlerFunc which sets content type of application/json and payload as a request body.
func SuccessJSON(payload interface{}) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(payload)
	})
}

// Error returns http.HandlerFunc which sets status code to provided value.
func Error(status int) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(status)
	})
}
