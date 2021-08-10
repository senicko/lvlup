// Package requests allows to make http requests in more convenient way.

package requests

import (
	"bytes"
	"io"
	"net/http"
)

type Requests struct {
	ApiBase    string
	HttpClient *http.Client
}

func NewRequests(apiBase string, httpClient *http.Client) *Requests {
	return &Requests{
		ApiBase:    apiBase,
		HttpClient: httpClient,
	}
}

type requestOptions struct {
	Headers map[string]string
	Query   map[string]string
	Body    io.Reader
}

type requestOption func(*requestOptions)

func WithHeaders(headers map[string]string) requestOption {
	return func(r *requestOptions) {
		r.Headers = headers
	}
}

func WithQuery(query map[string]string) requestOption {
	return func(r *requestOptions) {
		r.Query = query
	}
}

func WithBody(body []byte) requestOption {
	return func(r *requestOptions) {
		r.Body = bytes.NewBuffer(body)
	}
}

func newRequestOptions(opts ...requestOption) *requestOptions {
	requestOptions := &requestOptions{
		Body: http.NoBody,
	}

	for _, opt := range opts {
		opt(requestOptions)
	}

	return requestOptions
}

func (r Requests) Request(method string, path string, opts ...requestOption) (*http.Response, error) {
	requestOptions := newRequestOptions(opts...)

	request, err := http.NewRequest(method, r.ApiBase+path, requestOptions.Body)

	if err != nil {
		return nil, err
	}

	if requestOptions.Headers != nil {
		for key, value := range requestOptions.Headers {
			request.Header.Set(key, value)
		}
	}

	if requestOptions.Query != nil {
		query := request.URL.Query()

		for key, value := range requestOptions.Query {
			query.Set(key, value)
		}

		request.URL.RawQuery = query.Encode()
	}

	response, err := r.HttpClient.Do(request)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r Requests) Get(path string, opts ...requestOption) (*http.Response, error) {
	return r.Request(http.MethodGet, path, opts...)
}

func (r Requests) Post(path string, opts ...requestOption) (*http.Response, error) {
	return r.Request(http.MethodPost, path, opts...)
}
