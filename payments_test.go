package lvlup_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SeNicko/lvlup"
	"github.com/SeNicko/lvlup/internal/testutil"

	"github.com/stretchr/testify/assert"
)

func Test_Set_Redirect_Url_For_Payment(t *testing.T) {
	url := "url"
	options := &lvlup.CreatePaymentOptions{}

	lvlup.WithRedirect(url)(options)

	assert.Equal(t, options.RedirectUrl, url, "Redirect url set to invalid value")
}

func Test_Set_Webhook_Url_For_Payment(t *testing.T) {
	url := "url"
	options := &lvlup.CreatePaymentOptions{}

	lvlup.WithWebhook(url)(options)

	assert.Equal(t, options.WebhookUrl, url, "Webhook url set to invalid value")
}

func Test_Create_Payment(t *testing.T) {
	httpMock := httptest.NewServer(testutil.SuccessJSON(lvlup.CreatePaymentResult{}))
	defer httpMock.Close()
	client := lvlup.NewLvlClient("key", httpMock.Client())
	client.ApiBase = httpMock.URL

	_, err := client.CreatePayment("")

	assert.Nil(t, err, "Error should be nil")
}

func Test_Create_Payment_Error(t *testing.T) {
	httpMock := httptest.NewServer(testutil.Error(http.StatusInternalServerError))
	defer httpMock.Close()
	client := lvlup.NewLvlClient("", httpMock.Client())
	client.ApiBase = httpMock.URL

	_, err := client.CreatePayment("1.00")

	assert.NotNil(t, err, "Error should not be nil")
}

func Test_Set_Limit_For_List_Payments(t *testing.T) {
	limit := 1
	expected := "1"
	options := lvlup.ListPaymentsOptions{}

	lvlup.WithLimit(limit)(&options)

	assert.Equal(t, options["limit"], expected, "Limit set to invalid value")
}

func Test_Set_BeforeId_For_List_Payments(t *testing.T) {
	beforeId := 1
	expected := "1"
	options := lvlup.ListPaymentsOptions{}

	lvlup.WithBeforeId(beforeId)(&options)

	assert.Equal(t, options["beforeId"], expected, "BeforeId set to invalid value")
}

func Test_Set_AfterId_For_List_Payments(t *testing.T) {
	afterId := 1
	expected := "1"
	options := lvlup.ListPaymentsOptions{}

	lvlup.WithAfterId(afterId)(&options)

	assert.Equal(t, options["afterId"], expected, "AfterId set to invalid value")
}

func Test_List_Payments(t *testing.T) {
	httpMock := httptest.NewServer(testutil.SuccessJSON(lvlup.ListPaymentsResult{}))
	defer httpMock.Close()
	client := lvlup.NewLvlClient("", httpMock.Client())
	client.ApiBase = httpMock.URL

	_, err := client.ListPayments()

	assert.Nil(t, err, "Error should be nil")
}

func Test_List_Payments_Error(t *testing.T) {
	httpMock := httptest.NewServer(testutil.Error(http.StatusInternalServerError))
	defer httpMock.Close()
	client := lvlup.NewLvlClient("", httpMock.Client())
	client.ApiBase = httpMock.URL

	_, err := client.ListPayments()

	assert.NotNil(t, err, "Error should not be nil")
}

func Test_Get_Wallet_Balance(t *testing.T) {
	httpMock := httptest.NewServer(testutil.SuccessJSON(lvlup.WalletBalanceResult{}))
	defer httpMock.Close()
	client := lvlup.NewLvlClient("", httpMock.Client())
	client.ApiBase = httpMock.URL

	_, err := client.WalletBalance()

	assert.Nil(t, err, "Error should be nil")
}

func Test_Get_Wallet_Balance_Error(t *testing.T) {
	httpMock := httptest.NewServer(testutil.Error(http.StatusInternalServerError))
	defer httpMock.Close()
	client := lvlup.NewLvlClient("", httpMock.Client())
	client.ApiBase = httpMock.URL

	_, err := client.WalletBalance()

	assert.NotNil(t, err, "Error should not be nil")
}

func Test_Inspect_Payment(t *testing.T) {
	httpMock := httptest.NewServer(testutil.SuccessJSON(lvlup.InspectPaymentResult{}))
	defer httpMock.Close()
	client := lvlup.NewLvlClient("", httpMock.Client())
	client.ApiBase = httpMock.URL

	_, err := client.InspectPayment("")

	assert.Nil(t, err, "Error should be nil")
}

func Test_Inspect_Payment_Server_Error(t *testing.T) {
	httpMock := httptest.NewServer(testutil.Error(http.StatusInternalServerError))
	defer httpMock.Close()
	client := lvlup.NewLvlClient("", httpMock.Client())
	client.ApiBase = httpMock.URL

	_, err := client.InspectPayment("id")

	assert.NotNil(t, err, "Error should not be nil")
}

func Test_Inspect_Payment_NotFound_Error(t *testing.T) {
	assert := assert.New(t)
	httpMock := httptest.NewServer(testutil.Error(http.StatusNotFound))
	defer httpMock.Close()
	client := lvlup.NewLvlClient("", httpMock.Client())
	client.ApiBase = httpMock.URL

	result, err := client.InspectPayment("")

	assert.Nil(err, "Error should be nil")
	assert.Nil(result, "Result should be nil")
}
