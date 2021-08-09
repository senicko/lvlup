package lvlup_test

import (
	"testing"

	"github.com/SeNicko/lvlup"

	"github.com/stretchr/testify/assert"
)

func TestWithSandboxModeOption(t *testing.T) {
	assert := assert.New(t)

	testSandboxMode := true
	testSandboxApiBase := "https://api.lvlup.pro/v4/sandbox"
	testOptions := &lvlup.LvlClient{}

	lvlup.WithSandboxMode()(testOptions)

	assert.Equal(testOptions.ApiBase, testSandboxApiBase, "ApiBase set to sandbox api base")
	assert.Equal(testOptions.SandboxMode, testSandboxMode, "SandboxMode set to true")
}
