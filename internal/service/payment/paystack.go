package payment

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/rowjay007/event-bookie/config"
)

type PaystackClient struct {
    client *resty.Client
}

type PaystackPaymentResponse struct {
    Status  bool   `json:"status"`
    Message string `json:"message"`
    Data    struct {
        AuthorizationURL string `json:"authorization_url"`
        AccessCode       string `json:"access_code"`
        Reference        string `json:"reference"`
    } `json:"data"`
}

type PaystackVerificationResponse struct {
    Status  bool   `json:"status"`
    Message string `json:"message"`
    Data    struct {
        Amount    int64  `json:"amount"`
        Reference string `json:"reference"`
        Status    string `json:"status"`
    } `json:"data"`
}

func NewPaystackClient(config *config.Config) *PaystackClient {
    client := resty.New()
    client.SetHostURL("https://api.paystack.co")
    client.SetHeader("Authorization", "Bearer "+config.PaystackTestKey)
    return &PaystackClient{client: client}
}

func (p *PaystackClient) InitializePaystackPayment(amount int64, email string) (*PaystackPaymentResponse, error) {
    if amount <= 0 {
        return nil, errors.New("Invalid amount. Amount must be greater than zero")
    }

    amountInKobo := amount * 100
    payload := map[string]interface{}{
        "amount":   amountInKobo,
        "email":    email,
        "currency": "NGN",
    }

    resp, err := p.client.R().
        SetHeader("Content-Type", "application/json").
        SetBody(payload).
        Post("/transaction/initialize")

    if err != nil {
        return nil, err
    }

    var paystackPaymentResponse PaystackPaymentResponse
    err = json.Unmarshal(resp.Body(), &paystackPaymentResponse)
    if err != nil {
        return nil, err
    }

    paystackPaymentResponse.Data.Reference = "PST_" + paystackPaymentResponse.Data.Reference

    return &paystackPaymentResponse, nil
}


func (p *PaystackClient) VerifyPaystackPayment(reference string) (*PaystackVerificationResponse, error) {
    reference = strings.TrimPrefix(reference, "PST_")

    resp, err := p.client.R().
        Get(fmt.Sprintf("/transaction/verify/%s", reference))

    if err != nil {
        return nil, err
    }

    var verificationResponse PaystackVerificationResponse
    err = json.Unmarshal(resp.Body(), &verificationResponse)
    if err != nil {
        return nil, err
    }

    return &verificationResponse, nil
}

