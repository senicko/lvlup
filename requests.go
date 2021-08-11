package lvlup

import (
	"bytes"
	"io"
	"net/http"
)

// requestOptions represents options for http request.
type requestOptions struct {
	Headers map[string]string
	Query   map[string]string
	Body    io.Reader
}

type requestOption func(*requestOptions)

// withHeaders allows to set headers for a request.
func withHeaders(headers map[string]string) requestOption {
	return func(r *requestOptions) {
		r.Headers = headers
	}
}

// withQuery allows to set query for a request.
func withQuery(query map[string]string) requestOption {
	return func(r *requestOptions) {
		r.Query = query
	}
}

// withBody allows to set a body for a request.
func withBody(body []byte) requestOption {
	return func(r *requestOptions) {
		r.Body = bytes.NewBuffer(body)
	}
}

// newRequestOptions creates new requestOptions with applied settings.
func newRequestOptions(opts ...requestOption) *requestOptions {
	requestOptions := &requestOptions{
		Body: http.NoBody,
	}

	for _, opt := range opts {
		opt(requestOptions)
	}

	return requestOptions
}

// request allows to make a request to specified url.
// It returns recieved response and any errors encountered.
func (lc LvlClient) request(method string, path string, opts ...requestOption) (*http.Response, error) {
	requestOptions := newRequestOptions(opts...)

	request, err := http.NewRequest(method, lc.ApiBase+path, requestOptions.Body)

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

	response, err := lc.HttpClient.Do(request)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// get is a wrapper for request func.
// It sends get request to specified url.
func (lc LvlClient) get(path string, opts ...requestOption) (*http.Response, error) {
	return lc.request(http.MethodGet, path, opts...)
}

// post is a wrapper for request func.
// It sends get request to specified url.
func (lc LvlClient) post(path string, opts ...requestOption) (*http.Response, error) {
	return lc.request(http.MethodPost, path, opts...)
}
