package lvlup_test

import (
	"net/http"
	"testing"

	"github.com/SeNicko/lvlup"

	"github.com/stretchr/testify/assert"
)

func Test_create_client_with_sandbox_mode_option(t *testing.T) {
	client := lvlup.NewLvlClient("key", http.DefaultClient, lvlup.WithSandboxMode())

	assert.Equal(t, client.ApiBase, "https://api.sandbox.lvlup.pro/v4")
	assert.True(t, client.SandboxMode)
}
