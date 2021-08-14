package lvlup_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/SeNicko/lvlup"
	"github.com/SeNicko/lvlup/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func Test_List_All_Services(t *testing.T) {
	apiKey := "token"
	expectedApiKey := fmt.Sprintf("Bearer %v", apiKey)
	expectedPath := "/v4/services"

	handler := func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != expectedPath {
			return nil, fmt.Errorf("Request made to %v instead of %v", r.URL.Path, expectedPath)
		}

		if r.Header.Get("Authorization") != expectedApiKey {
			return nil, fmt.Errorf("Invalid authorization token format")
		}

		rBody, err := json.Marshal(&lvlup.ListServicesResult{})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(rBody)),
		}, nil
	}

	client := testutil.NewTestLvlClient(apiKey, handler)

	_, err := client.ListServices()

	assert.Nil(t, err, "Error should be nil")
}

func Test_List_All_Services_Server_Error(t *testing.T) {
	client := testutil.NewTestLvlClient("token", testutil.HttpError(http.StatusInternalServerError))

	_, err := client.ListServices()

	assert.NotNil(t, err, "Error should not be nil")
}

func Test_List_DDoS_Attacks(t *testing.T) {
	vpsId := "1"
	apiKey := "token"
	expectedPath := fmt.Sprintf("/v4/services/vps/%v/attacks", vpsId)
	expectedApiKey := fmt.Sprintf("Bearer %v", apiKey)

	handler := func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != expectedPath {
			return nil, fmt.Errorf("Request made to %v instead of %v", r.URL.Path, expectedPath)
		}

		if r.Header.Get("Authorization") != expectedApiKey {
			return nil, fmt.Errorf("Invalid authorization token format")
		}

		rBody, err := json.Marshal(lvlup.ListDDoSAttacksResult{})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(rBody)),
		}, nil
	}

	client := testutil.NewTestLvlClient("token", handler)

	_, err := client.ListDDoSAttacks(vpsId)

	assert.Nil(t, err, "Error should be nil")
}

func Test_List_DDoS_Attacks_Server_Error(t *testing.T) {
	client := testutil.NewTestLvlClient("token", testutil.HttpError(http.StatusBadRequest))

	_, err := client.ListDDoSAttacks("1")

	assert.NotNil(t, err, "Error should not be nil")
}

func Test_Get_UDP_Filter(t *testing.T) {
	vpsId := "1"
	apiKey := "token"
	expectedPath := fmt.Sprintf("/v4/services/vps/%v/filtering", vpsId)
	expectedApiKey := fmt.Sprintf("Bearer %v", apiKey)

	handler := func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != expectedPath {
			return nil, fmt.Errorf("Request made to %v instead of %v", r.URL.Path, expectedPath)
		}

		if r.Header.Get("Authorization") != expectedApiKey {
			return nil, fmt.Errorf("Invalid authorization token format")
		}

		rBody, err := json.Marshal(lvlup.GetUDPFilterResult{})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(rBody)),
		}, nil
	}

	client := testutil.NewTestLvlClient("token", handler)

	_, err := client.GetUDPFilter(vpsId)

	assert.Nil(t, err, "Error should be nil")
}

func Test_Get_UDP_Filter_Server_Error(t *testing.T) {
	client := testutil.NewTestLvlClient("token", testutil.HttpError(http.StatusInternalServerError))

	_, err := client.GetUDPFilter("1")

	assert.NotNil(t, err, "Error should not be nil")
}

func Test_Set_UDP_Filtering(t *testing.T) {
	vpsId := "1"
	apiKey := "token"
	filteringEnabled := true
	expectedPath := fmt.Sprintf("/v4/services/vps/%v/filtering", vpsId)
	expectedApiKey := fmt.Sprintf("Bearer %v", apiKey)

	handler := func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != expectedPath {
			return nil, fmt.Errorf("Request made to %v instead of %v", r.URL.Path, expectedPath)
		}

		if r.Header.Get("Authorization") != expectedApiKey {
			return nil, fmt.Errorf("Invalid authorization token format")
		}

		var body lvlup.SetUDPFilteringOptions
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			return nil, err
		}

		if body.FilteringEnabled != filteringEnabled {
			return nil, fmt.Errorf("FilteringEnabled set to %v instead of %v", body.FilteringEnabled, filteringEnabled)
		}

		rBody, err := json.Marshal(lvlup.SetUDPFilteringResult{})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(rBody)),
		}, nil
	}

	client := testutil.NewTestLvlClient("token", handler)

	_, err := client.SetUDPFiltering(vpsId, filteringEnabled)

	assert.Nil(t, err, "Error should be nil")
}

func Test_Set_UDP_Filtering_Server_Error(t *testing.T) {
	client := testutil.NewTestLvlClient("token", testutil.HttpError(http.StatusInternalServerError))

	_, err := client.SetUDPFiltering("1", true)

	assert.NotNil(t, err, "Error should not be nil")
}

func Test_List_UDP_Filter_Exceptions(t *testing.T) {
	vpsId := "1"
	apiKey := "token"
	expectedPath := fmt.Sprintf("/v4/services/vps/%v/filtering/whitelist", vpsId)
	expectedApiKey := fmt.Sprintf("Bearer %v", apiKey)

	handler := func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != expectedPath {
			return nil, fmt.Errorf("Request made to %v instead of %v", r.URL.Path, expectedPath)
		}

		if r.Header.Get("Authorization") != expectedApiKey {
			return nil, fmt.Errorf("Invalid authorization token format")
		}

		rBody, err := json.Marshal([]lvlup.UDPFilterException{})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(rBody)),
		}, nil
	}

	client := testutil.NewTestLvlClient(apiKey, handler)

	_, err := client.ListUDPFilterExceptions(vpsId)

	assert.Nil(t, err, "Error should be nil")
}

func Test_List_UDP_Filter_Exceptions_Server_Error(t *testing.T) {
	client := testutil.NewTestLvlClient("token", testutil.HttpError(http.StatusInternalServerError))

	_, err := client.ListUDPFilterExceptions("1")

	assert.NotNil(t, err, "Error should not be nil")
}

func Test_Add_UDP_Filter_Exception(t *testing.T) {
	vpsId := "1"
	apiKey := "token"
	expectedPath := fmt.Sprintf("/v4/services/vps/%v/filtering/whitelist", vpsId)
	expectedApiKey := fmt.Sprintf("Bearer %v", apiKey)

	handler := func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != expectedPath {
			return nil, fmt.Errorf("Request made to %v instead of %v", r.URL.Path, expectedPath)
		}

		if r.Header.Get("Authorization") != expectedApiKey {
			return nil, fmt.Errorf("Invalid authorization token format")
		}

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	}

	client := testutil.NewTestLvlClient(apiKey, handler)

	err := client.AddUDPFilterException(vpsId, &lvlup.UDPFilterException{})

	assert.Nil(t, err, "Error should be nil")
}

func Test_Add_UDP_Filter_Exception_Server_Error(t *testing.T) {
	client := testutil.NewTestLvlClient("token", testutil.HttpError(http.StatusInternalServerError))

	err := client.AddUDPFilterException("1", &lvlup.UDPFilterException{})

	assert.NotNil(t, err, "Error should not be nil")
}

func Test_Remove_UDP_Filter_Exception(t *testing.T) {
	vpsId := "1"
	apiKey := "token"
	exceptionId := "1"
	expectedPath := fmt.Sprintf("/v4/services/vps/%v/filtering/whitelist/%v", vpsId, exceptionId)
	expectedApiKey := fmt.Sprintf("Bearer %v", apiKey)

	handler := func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != expectedPath {
			return nil, fmt.Errorf("Request made to %v instead of %v", r.URL.Path, expectedPath)
		}

		if r.Header.Get("Authorization") != expectedApiKey {
			return nil, fmt.Errorf("Invalid authorization token format")
		}

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	}

	client := testutil.NewTestLvlClient(apiKey, handler)

	err := client.RemoveUDPFilterException(vpsId, exceptionId)

	assert.Nil(t, err, "Error should be nil")
}

func Test_Remove_UDP_Filter_Exception_Server_Error(t *testing.T) {
	client := testutil.NewTestLvlClient("token", testutil.HttpError(http.StatusInternalServerError))

	err := client.RemoveUDPFilterException("1", "1")

	assert.NotNil(t, err, "Error should not be nil")
}

func Test_Get_Proxmo_User(t *testing.T) {
	vpsId := "1"
	apiKey := "token"
	expectedPath := fmt.Sprintf("/v4/services/vps/%v/proxmo", vpsId)
	expectedApiKey := fmt.Sprintf("Bearer %v", apiKey)

	handler := func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != expectedPath {
			return nil, fmt.Errorf("Request made to %v instead of %v", r.URL.Path, expectedPath)
		}

		if r.Header.Get("Authorization") != expectedApiKey {
			return nil, fmt.Errorf("Invalid authorization token format")
		}

		rBody, err := json.Marshal(lvlup.ProxmoUser{})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(rBody)),
		}, nil
	}

	client := testutil.NewTestLvlClient(apiKey, handler)

	_, err := client.GetProxmoUser(vpsId)

	assert.Nil(t, err, "Error should be nil")
}

func Test_Get_Proxmo_User_Error(t *testing.T) {
	client := testutil.NewTestLvlClient("token", testutil.HttpError(http.StatusInternalServerError))

	_, err := client.GetProxmoUser("1")

	assert.NotNil(t, err, "Error should not be nil")
}

func Test_Start_VPS(t *testing.T) {
	vpsId := "1"
	apiKey := "token"
	expectedPath := fmt.Sprintf("/v4/services/vps/%v/start", vpsId)
	expectedApiKey := fmt.Sprintf("Bearer %v", apiKey)

	handler := func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != expectedPath {
			return nil, fmt.Errorf("Request made to %v instead of %v", r.URL.Path, expectedApiKey)
		}

		if r.Header.Get("Authorization") != expectedApiKey {
			return nil, fmt.Errorf("Invalid authorization token format")
		}

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	}

	client := testutil.NewTestLvlClient(apiKey, handler)

	err := client.StartVPS(vpsId)

	assert.Nil(t, err, "Error should be nil")
}

func Test_Start_VPS_Error(t *testing.T) {
	client := testutil.NewTestLvlClient("token", testutil.HttpError(http.StatusInternalServerError))

	err := client.StartVPS("1")

	assert.NotNil(t, err, "Error should not be nil")
}

func Test_Get_VPS_State(t *testing.T) {
	vpsId := "1"
	apiKey := "token"
	expectedPath := fmt.Sprintf("/v4/services/vps/%v/state", vpsId)
	expectedApiKey := fmt.Sprintf("Bearer %s", apiKey)

	handler := func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != expectedPath {
			return nil, fmt.Errorf("Request made to %v instead of %v", r.URL.Path, expectedApiKey)
		}

		if r.Header.Get("Authorization") != expectedApiKey {
			return nil, fmt.Errorf("Invalid authorization token format")
		}

		rBody, err := json.Marshal(lvlup.GetVPSStateResult{})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(rBody)),
		}, nil
	}

	client := testutil.NewTestLvlClient(apiKey, handler)

	_, err := client.GetVPSState(vpsId)

	assert.Nil(t, err, "Error should be nil")
}

func Test_Get_VPS_State_Server_Error(t *testing.T) {
	client := testutil.NewTestLvlClient("token", testutil.HttpError(http.StatusInternalServerError))

	_, err := client.GetVPSState("1")

	assert.NotNil(t, err, "Error should not be nil")
}

func Test_Stop_VPS(t *testing.T) {
	vpsId := "1"
	apiKey := "token"
	expectedPath := fmt.Sprintf("/v4/services/vps/%v/stop", vpsId)
	expectedApiKey := fmt.Sprintf("Bearer %s", apiKey)

	handler := func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != expectedPath {
			return nil, fmt.Errorf("Request made to %v instead of %v", r.URL.Path, expectedPath)
		}

		if r.Header.Get("Authorization") != expectedApiKey {
			return nil, fmt.Errorf("Invalid authorization token format")
		}

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	}

	client := testutil.NewTestLvlClient(apiKey, handler)

	err := client.StopVPS(vpsId)

	assert.Nil(t, err, "Error should be nil")
}

func Test_Stop_VPS_Error(t *testing.T) {
	client := testutil.NewTestLvlClient("token", testutil.HttpError(http.StatusInternalServerError))

	err := client.StopVPS("1")

	assert.NotNil(t, err, "Error should not be nil")
}