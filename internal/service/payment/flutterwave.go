package payment

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/rowjay007/event-bookie/config"
)

type FlutterwaveClient struct {
    client *resty.Client
}

type FlutterwavePaymentResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Data    struct {
        Link string `json:"link"`
    } `json:"data"`
}

type FlutterwavePaymentVerificationResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Data    struct {
        Amount    int64  `json:"amount"`
        Currency  string `json:"currency"`
        Status    string `json:"status"`
        Reference string `json:"tx_ref"`
    } `json:"data"`
}

func NewFlutterwaveClient(config *config.Config) *FlutterwaveClient {
    client := resty.New()
    client.SetHostURL("https://api.flutterwave.com/v3")
    client.SetHeader("Authorization", "Bearer "+config.FlutterwaveTestKey)
    return &FlutterwaveClient{client: client}
}

func (f *FlutterwaveClient) InitializePayment(amount int64, email, txRef string) (*FlutterwavePaymentResponse, error) {
    if amount <= 0 {
        return nil, errors.New("invalid amount")
    }

    payload := map[string]interface{}{
        "tx_ref":        txRef,
        "amount":        amount,
        "currency":      "NGN",
        "redirect_url":  "http://your-redirect-url.com",
        "payment_type":  "card",
        "customer": map[string]string{"email": email},
    }

    resp, err := f.client.R().
        SetHeader("Content-Type", "application/json").
        SetBody(payload).
        Post("/payments")

    if err != nil {
        return nil, err
    }

    var paymentResponse FlutterwavePaymentResponse
    err = json.Unmarshal(resp.Body(), &paymentResponse)
    if err != nil {
        return nil, err
    }

    return &paymentResponse, nil
}

func (f *FlutterwaveClient) VerifyPayment(txRef string) (*FlutterwavePaymentVerificationResponse, error) {
    resp, err := f.client.R().
        Get(fmt.Sprintf("/transactions/%s/verify", txRef))

    if err != nil {
        return nil, err
    }

    var verificationResponse FlutterwavePaymentVerificationResponse
    err = json.Unmarshal(resp.Body(), &verificationResponse)
    if err != nil {
        return nil, err
    }

    return &verificationResponse, nil
}
