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

func Test_set_redirect_url_for_payment(t *testing.T) {
	redirectUrl := "url"

	handler := func(r *http.Request) (*http.Response, error) {
		var body lvlup.CreatePaymentOptions
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			return nil, err
		}

		if body.RedirectUrl != redirectUrl {
			return nil, fmt.Errorf("RedirectUrl set to %v instead of %v", body.RedirectUrl, redirectUrl)
		}

		rBody, err := json.Marshal(lvlup.CreatePaymentResult{})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(rBody)),
		}, nil
	}

	client := testutil.NewTestLvlClient("token", handler)

	_, err := client.CreatePayment("1.00", lvlup.WithRedirect(redirectUrl))

	assert.Nil(t, err, "Error should be nil")
}

func Test_set_webhook_url_for_payment(t *testing.T) {
	webhookUrl := "url"

	handler := func(r *http.Request) (*http.Response, error) {
		var body lvlup.CreatePaymentOptions
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			return nil, err
		}

		if body.WebhookUrl != webhookUrl {
			return nil, fmt.Errorf("RedirectUrl set to %v instead of %v", body.WebhookUrl, webhookUrl)
		}

		rBody, err := json.Marshal(lvlup.CreatePaymentResult{})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(rBody)),
		}, nil
	}

	client := testutil.NewTestLvlClient("token", handler)

	_, err := client.CreatePayment("1.00", lvlup.WithWebhook(webhookUrl))

	assert.Nil(t, err, "Error should be nil")
}

func Test_create_payment(t *testing.T) {
	apiKey := "token"
	amount := "10.00"
	expectedPath := "/v4/wallet/up"
	expectedApiKey := "Bearer " + apiKey

	handler := func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != expectedPath {
			return nil, fmt.Errorf("Request made to %v instead of %v", r.URL.Path, expectedPath)
		}

		if r.Header.Get("Authorization") != expectedApiKey {
			return nil, fmt.Errorf("Invalid authorization token format")
		}

		var body lvlup.CreatePaymentOptions
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			return nil, err
		}

		if body.Amount != amount {
			return nil, fmt.Errorf("Amount set to %v instead of %v", body.Amount, amount)
		}

		rBody, err := json.Marshal(lvlup.CreatePaymentResult{})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(rBody)),
		}, nil
	}

	client := testutil.NewTestLvlClient(apiKey, handler)

	_, err := client.CreatePayment(amount)

	assert.Nil(t, err, "Error should be nil")
}

func Test_create_payment_server_error(t *testing.T) {
	client := testutil.NewTestLvlClient("token", testutil.HttpError(http.StatusInternalServerError))

	_, err := client.CreatePayment("1.00")

	assert.NotNil(t, err, "Error should not be nil")
}

func Test_set_limit_for_list_payments(t *testing.T) {
	testLimit := 10
	expectedLimit := "10"

	handler := func(r *http.Request) (*http.Response, error) {
		limit := r.URL.Query().Get("limit")

		if limit != expectedLimit {
			return nil, fmt.Errorf("AfterId set to %v instead of %v", limit, expectedLimit)
		}

		rBody, err := json.Marshal(lvlup.ListPaymentsResult{})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(rBody)),
		}, nil
	}

	client := testutil.NewTestLvlClient("token", handler)

	_, err := client.ListPayments(lvlup.WithLimit(testLimit))

	assert.Nil(t, err, "Error should be nil")
}

func Test_set_beforeId_for_list_payments(t *testing.T) {
	testBeforeId := 10
	expectedBeforeId := "10"

	handler := func(r *http.Request) (*http.Response, error) {
		beforeId := r.URL.Query().Get("beforeId")

		if beforeId != expectedBeforeId {
			return nil, fmt.Errorf("AfterId set to %v instead of %v", beforeId, expectedBeforeId)
		}

		rBody, err := json.Marshal(lvlup.ListPaymentsResult{})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(rBody)),
		}, nil
	}

	client := testutil.NewTestLvlClient("token", handler)

	_, err := client.ListPayments(lvlup.WithBeforeId(testBeforeId))

	assert.Nil(t, err, "Error should be nil")
}

func Test_set_afterId_for_list_payments(t *testing.T) {
	testAfterId := 10
	expectedAfterId := "10"

	handler := func(r *http.Request) (*http.Response, error) {
		afterId := r.URL.Query().Get("afterId")

		if afterId != expectedAfterId {
			return nil, fmt.Errorf("AfterId set to %v instead of %v", afterId, expectedAfterId)
		}

		rBody, err := json.Marshal(lvlup.ListPaymentsResult{})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(rBody)),
		}, nil
	}

	client := testutil.NewTestLvlClient("token", handler)

	_, err := client.ListPayments(lvlup.WithAfterId(testAfterId))

	assert.Nil(t, err, "Error should be nil")
}

func Test_list_payments(t *testing.T) {
	apiKey := "token"
	expectedPath := "/v4/payments"
	expectedApiKey := "Bearer " + apiKey

	handler := func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != expectedPath {
			return nil, fmt.Errorf("Request made to %v instead of %v", r.URL.Path, expectedPath)
		}

		if r.Header.Get("Authorization") != expectedApiKey {
			return nil, fmt.Errorf("Invalid authorization token format")
		}

		rBody, err := json.Marshal(lvlup.ListPaymentsResult{})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(rBody)),
		}, nil
	}

	client := testutil.NewTestLvlClient(apiKey, handler)

	_, err := client.ListPayments()

	assert.Nil(t, err, "Error should be nil")
}

func Test_list_payments_server_error(t *testing.T) {
	client := testutil.NewTestLvlClient("token", testutil.HttpError(http.StatusInternalServerError))

	_, err := client.ListPayments()

	assert.NotNil(t, err, "Error should not be nil")
}

func Test_get_wallet_balance(t *testing.T) {
	apiKey := "token"
	expectedPath := "/v4/wallet"
	expectedApiKey := "Bearer " + apiKey

	handler := func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != expectedPath {
			return nil, fmt.Errorf("Request made to %v instead of %v", r.URL.Path, expectedPath)
		}

		if r.Header.Get("Authorization") != expectedApiKey {
			return nil, fmt.Errorf("Invalid authorization token format")
		}

		rBody, err := json.Marshal(lvlup.WalletBalanceResult{})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(rBody)),
		}, nil
	}

	client := testutil.NewTestLvlClient(apiKey, handler)

	_, err := client.WalletBalance()

	assert.Nil(t, err, "Error should be nil")
}

func Test_get_wallet_balance_server_error(t *testing.T) {
	client := testutil.NewTestLvlClient("token", testutil.HttpError(http.StatusInternalServerError))

	_, err := client.WalletBalance()

	assert.NotNil(t, err, "Error should not be nil")
}

func Test_inspect_payment(t *testing.T) {
	apiKey := "token"
	paymentId := "1"
	expectedPath := "/v4/wallet/up/" + paymentId
	expectedApiKey := "Bearer " + apiKey

	handler := func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != expectedPath {
			return nil, fmt.Errorf("Request made to %v instead of %v", r.URL.Path, expectedPath)
		}

		if r.Header.Get("Authorization") != expectedApiKey {
			return nil, fmt.Errorf("Invalid authorization token format")
		}

		rBody, err := json.Marshal(lvlup.InspectPaymentResult{})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(rBody)),
		}, nil
	}

	client := testutil.NewTestLvlClient(apiKey, handler)

	_, err := client.InspectPayment(paymentId)

	assert.Nil(t, err, "Error should be nil")
}

func Test_inspect_payment_server_error(t *testing.T) {
	client := testutil.NewTestLvlClient("token", testutil.HttpError(http.StatusInternalServerError))

	_, err := client.InspectPayment("id")

	assert.NotNil(t, err, "Error should not be nil")
}

func Test_inspect_payment_not_found_error(t *testing.T) {
	client := testutil.NewTestLvlClient("token", testutil.HttpError(http.StatusNotFound))

	result, err := client.InspectPayment("")

	assert.Nil(t, err, "Error should be nil")
	assert.Nil(t, result, "Result should be nil")
}
