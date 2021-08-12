package lvlup_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SeNicko/lvlup"
	"github.com/SeNicko/lvlup/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_List_All_Services(t *testing.T) {
	httpClient := httptest.NewServer(testutil.SuccessJSON(lvlup.ListServicesResult{}))
	defer httpClient.Close()
	client := lvlup.NewLvlClient("", httpClient.Client())
	client.ApiBase = httpClient.URL

	_, err := client.ListServices()

	assert.Nil(t, err, "Error should be nil")
}

func Test_List_All_Services_Error(t *testing.T) {
	httpClient := httptest.NewServer(testutil.Error(http.StatusInternalServerError))
	defer httpClient.Close()
	client := lvlup.NewLvlClient("", httpClient.Client())
	client.ApiBase = httpClient.URL

	_, err := client.ListServices()

	assert.NotNil(t, err, "Error should not be nil")
}