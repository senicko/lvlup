package lvlup_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SeNicko/lvlup"

	"github.com/stretchr/testify/assert"
)

func TestCreatePaymentWithRedirectOption(t *testing.T) {
	testRedirectUrl := "url"
	testOptions := &lvlup.CreatePaymentOptions{}

	lvlup.WithRedirect(testRedirectUrl)(testOptions)

	assert.Equal(t, testOptions.RedirectUrl, testRedirectUrl, "RedirectUrl should be set")
}

func TestCreatePaymentWithWebhookUrl(t *testing.T) {
	testWebhookUrl := "url"
	testOptions := &lvlup.CreatePaymentOptions{}

	lvlup.WithWebhook(testWebhookUrl)(testOptions)

	assert.Equal(t, testOptions.WebhookUrl, testWebhookUrl, "WebhookUrl should be set")
}

func successfullResponse(payload interface{}) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(payload)
	}
}

func errorResponse(code int) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(code)
	}
}

func TestCreatePayment(t *testing.T) {
	assert := assert.New(t)

	testResponse := &lvlup.CreatePaymentResult{
		Id:  "id",
		Url: "url",
	}

	httpMock := httptest.NewServer(http.HandlerFunc(successfullResponse(testResponse)))
	defer httpMock.Close()

	client := lvlup.NewLvlClient("key", httpMock.Client())
	client.ApiBase = httpMock.URL

	result, err := client.CreatePayment("9.99")

	assert.Nil(err, "Error should be nil")
	assert.NotNil(result, "Resoul should be not nil")
}

func TestCreatePaymentError(t *testing.T) {
	assert := assert.New(t)

	httpMock := httptest.NewServer(http.HandlerFunc(errorResponse(http.StatusBadRequest)))
	defer httpMock.Close()

	client := lvlup.NewLvlClient("key", httpMock.Client())
	client.ApiBase = httpMock.URL

	result, err := client.CreatePayment("9.99")

	assert.NotNil(err, "Error should be not nil")
	assert.Nil(result, "Result should be nil")
}

func TestListPaymentsWithLimitOption(t *testing.T) {
	testLimit := 20
	testOptions := lvlup.ListPaymentsOptions{}

	lvlup.WithLimit(testLimit)(&testOptions)

	assert.Equal(t, testOptions["limit"], testLimit, "Limit should be equal 20")
}

func TestListPaymentsWithBeforeIdOption(t *testing.T) {
	testBeforeId := 20
	testOptions := lvlup.ListPaymentsOptions{}

	lvlup.WithBeforeId(testBeforeId)(&testOptions)

	assert.Equal(t, testOptions["beforeId"], testBeforeId, "BeforeId should be set to 20")
}

func TestListPaymentsWithAfterIdOption(t *testing.T) {
	testAfterId := 20
	testOptions := lvlup.ListPaymentsOptions{}

	lvlup.WithAfterId(testAfterId)(&testOptions)

	assert.Equal(t, testOptions["afterId"], testAfterId, "AfterId should be set to 20")
}

func TestListPayments(t *testing.T) {
	assert := assert.New(t)

	testResponse := &lvlup.ListPaymentsResult{
		Count: 0,
		Items: []lvlup.ListPaymentsResultItem{},
	}

	httpMock := httptest.NewServer(http.HandlerFunc(successfullResponse(testResponse)))
	defer httpMock.Close()

	client := lvlup.NewLvlClient("key", httpMock.Client())
	client.ApiBase = httpMock.URL

	result, err := client.ListPayments()

	assert.Nil(err, "Error should be nil")
	assert.NotNil(result, "Result should be not nil")
}

func TestListPaymentsError(t *testing.T) {
	assert := assert.New(t)

	httpMock := httptest.NewServer(http.HandlerFunc(errorResponse(http.StatusInternalServerError)))
	defer httpMock.Close()

	client := lvlup.NewLvlClient("key", httpMock.Client())
	client.ApiBase = httpMock.URL

	result, err := client.ListPayments()

	assert.NotNil(err, "Error should not be nil")
	assert.Nil(result, "Result should be nil")
}

func TestWalletBalance(t *testing.T) {
	assert := assert.New(t)

	testResult := &lvlup.WalletBalanceResult{
		BalancePlnFormatted: "0PLN :P",
		BalancePlnInt:       0,
	}

	httpMock := httptest.NewServer(http.HandlerFunc(successfullResponse(testResult)))
	defer httpMock.Close()

	client := lvlup.NewLvlClient("key", httpMock.Client())
	client.ApiBase = httpMock.URL

	result, err := client.WalletBalance()

	assert.Nil(err, "Error should be nil")
	assert.NotNil(result, "Result should not be nil")
}

func TestWalletBalanceError(t *testing.T) {
	assert := assert.New(t)

	httpMock := httptest.NewServer(http.HandlerFunc(errorResponse(http.StatusInternalServerError)))
	defer httpMock.Close()

	client := lvlup.NewLvlClient("key", httpMock.Client())
	client.ApiBase = httpMock.URL

	result, err := client.WalletBalance()

	assert.NotNil(err, "Error should not be nil")
	assert.Nil(result, "Result should be nil")
}

func TestInspectPayment(t *testing.T) {
	assert := assert.New(t)

	testResult := &lvlup.InspectPaymentResult{
		AmountInt:        1,
		AmountStr:        "1",
		AmountWithFeeInt: 2,
		AmountWithFeeStr: "2",
		Payed:            true,
	}

	httpMock := httptest.NewServer(http.HandlerFunc(successfullResponse(testResult)))
	defer httpMock.Close()

	client := lvlup.NewLvlClient("key", httpMock.Client())
	client.ApiBase = httpMock.URL

	result, err := client.InspectPayment("id")

	assert.Nil(err, "Error should be nil")
	assert.NotNil(result, "Result should not be nil")
}

func TestInspectPaymentError(t *testing.T) {
	assert := assert.New(t)

	httpMock := httptest.NewServer(http.HandlerFunc(errorResponse(http.StatusInternalServerError)))
	defer httpMock.Close()

	client := lvlup.NewLvlClient("key", httpMock.Client())
	client.ApiBase = httpMock.URL

	result, err := client.InspectPayment("id")

	assert.NotNil(err, "Error should not be nil")
	assert.Nil(result, "Result should be nil")
}

func TestInspectPaymentNotFound(t *testing.T) {
	assert := assert.New(t)

	httpMock := httptest.NewServer(http.HandlerFunc(errorResponse(http.StatusNotFound)))
	defer httpMock.Close()

	client := lvlup.NewLvlClient("key", httpMock.Client())
	client.ApiBase = httpMock.URL

	result, err := client.InspectPayment("id")

	assert.Nil(err, "Error should be nil")
	assert.Nil(result, "Result should be nil")
}
