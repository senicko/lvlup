package lvlup

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// CreatePaymentOptions represents available options for CreatePayment func.
type CreatePaymentOptions struct {
	Amount      string `json:"amount"`
	RedirectUrl string `json:"redirectUrl"`
	WebhookUrl  string `json:"webhookUrl"`
}

// CreatePaymentResult represents result of CreatePayment func.
type CreatePaymentResult struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

// CreatePaymentOption represents a functional option for CreatePayment func.
type CreatePaymentOption func(*CreatePaymentOptions)

// Set url to which user should be redirected after completing a payment.
func WithRedirect(url string) CreatePaymentOption {
	return func(cpo *CreatePaymentOptions) {
		cpo.RedirectUrl = url
	}
}

// WithWebhook sets webhook url to which POST request will be send after completing a payment.
func WithWebhook(url string) CreatePaymentOption {
	return func(cpo *CreatePaymentOptions) {
		cpo.WebhookUrl = url
	}
}

// CreatePayment allows to create a new payment url.
// It returns result of a request and any errors encountered.
func (lc LvlClient) CreatePayment(amount string, opts ...CreatePaymentOption) (*CreatePaymentResult, error) {
	options := &CreatePaymentOptions{
		Amount:      amount,
		RedirectUrl: "",
		WebhookUrl:  "",
	}

	for _, opt := range opts {
		opt(options)
	}

	payload, err := json.Marshal(options)

	if err != nil {
		return nil, err
	}

	response, err := lc.post(
		"/wallet/up",
		withBody(payload),
		withHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", lc.ApiKey),
		}),
	)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(response.Status)
	}

	var result CreatePaymentResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ListPaymentsResultItem represents single item from ListPayments func.
type ListPaymentsResultItem struct {
	Amount      string `json:"amount"`
	CreatedAt   string `json:"createdAt"`
	Description string `json:"description"`
	Id          int    `json:"id"`
	MethodId    int    `json:"methodId"`
	ServiceId   int    `json:"serviceId"`
}

// ListPaymentsResult represents result of ListPayments func.
type ListPaymentsResult struct {
	Count int                      `json:"count"`
	Items []ListPaymentsResultItem `json:"items"`
}

// ListPaymentsOptions represents a map storing optional query parameters for ListPayments func.
type ListPaymentsOptions map[string]string

// ListPaymentsOption represents functional option for a ListPayments func.
type ListPaymentsOption func(*ListPaymentsOptions)

// WithLimit allows to set max payments count per page.
func WithLimit(limit int) ListPaymentsOption {
	return func(lpo *ListPaymentsOptions) {
		(*lpo)["limit"] = strconv.Itoa(limit)
	}
}

// WithBeforeId allows to set payment id before which payments should be returned.
func WithBeforeId(beforeId int) ListPaymentsOption {
	return func(lpo *ListPaymentsOptions) {
		(*lpo)["beforeId"] = strconv.Itoa(beforeId)
	}
}

// WithAfterId allows to set payment id after which payments should be returned.
func WithAfterId(afterId int) ListPaymentsOption {
	return func(lpo *ListPaymentsOptions) {
		(*lpo)["afterId"] = strconv.Itoa(afterId)
	}
}

// ListPayments allows to list client's payments.
// It returns request result and eny errors encountered.
func (lc LvlClient) ListPayments(opts ...ListPaymentsOption) (*ListPaymentsResult, error) {
	var options ListPaymentsOptions = map[string]string{}

	for _, opt := range opts {
		opt(&options)
	}

	response, err := lc.get(
		"/payments",
		withQuery(options),
		withHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", lc.ApiKey),
		}),
	)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(response.Status)
	}

	var result ListPaymentsResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// WalletBalanceResult represents result of WalletBalance request.
type WalletBalanceResult struct {
	BalancePlnFormatted string `json:"balancePlnFormatted"`
	BalancePlnInt       int    `json:"balancePlnInt"`
}

// WalletBalance allows to get current wallet balance.
// It returns request result and any errors encountered.
func (lc LvlClient) WalletBalance() (*WalletBalanceResult, error) {
	response, err := lc.get(
		"/wallet",
		withHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", lc.ApiKey),
		}),
	)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(response.Status)
	}

	var result WalletBalanceResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// InspectPaymentResult represents result of InspectPayment func.
type InspectPaymentResult struct {
	AmountInt        int    `json:"amountInt"`
	AmountStr        string `json:"amuntStr"`
	AmountWithFeeInt int    `json:"amountWithFeeInt"`
	AmountWithFeeStr string `json:"amountWithFeeStr"`
	Payed            bool   `json:"payed"`
}

// InspectPayment allows to inspect a payment.
// It returns request result and any errors encountered.
func (lc LvlClient) InspectPayment(paymentId string) (*InspectPaymentResult, error) {
	response, err := lc.get(
		"/wallet/up/"+paymentId,
		withHeaders(map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", lc.ApiKey),
		}),
	)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return nil, nil
	} else if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(response.Status)
	}

	var result InspectPaymentResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
