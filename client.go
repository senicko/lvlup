package lvlup

import (
	"net/http"
)

// LvlClient describes properties stored by the client.
type LvlClient struct {
	ApiKey      string
	ApiBase     string
	SandboxMode bool
	HttpClient  *http.Client
}

// LvlClientOption describes functional option for the client.
type LvlClientOption func(*LvlClient)

// WithSandboxMode enables sandbox mode for the LvlClient.
func WithSandboxMode() LvlClientOption {
	return func(lc *LvlClient) {
		lc.SandboxMode = true
		lc.ApiBase = "https://api.sandbox.lvlup.pro/v4"
	}
}

// NewLvlClient creates new lvlup api client.
func NewLvlClient(apiKey string, httpClient *http.Client, opts ...LvlClientOption) *LvlClient {
	lc := &LvlClient{
		ApiKey:      apiKey,
		ApiBase:     "https://api.lvlup.pro/v4",
		SandboxMode: false,
		HttpClient:  httpClient,
	}

	for _, opt := range opts {
		opt(lc)
	}

	return lc
}
