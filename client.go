package lvlup

import (
	"net/http"

	"github.com/SeNicko/lvlup/requests"
)

// LvlClient describes properties stored by the client
type LvlClient struct {
	ApiKey      string
	ApiBase     string
	SandboxMode bool
	Requests    *requests.Requests
}

// LvlClientOption describes functional option for the client
type LvlClientOption func(*LvlClient)

// EnableSandboxMode enables sandbox mode for the LvlClient
func WithSandboxMode() LvlClientOption {
	return func(lc *LvlClient) {
		lc.SandboxMode = true
		lc.ApiBase = "https://api.lvlup.pro/v4/sandbox"
	}
}

// NewLvlClient creates new lvlup api client
func NewLvlClient(apiKey string, httpClient *http.Client, opts ...LvlClientOption) *LvlClient {
	requests := requests.NewRequests("", httpClient)

	lc := &LvlClient{
		ApiKey:      apiKey,
		ApiBase:     "https://api.lvlup.pro/v4",
		SandboxMode: false,
		Requests:    requests,
	}

	for _, opt := range opts {
		opt(lc)
	}

	return lc
}
