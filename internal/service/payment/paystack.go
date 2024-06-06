package payment

import (
    "encoding/json"
    "errors"
    "fmt"

    "github.com/go-resty/resty/v2"
    "github.com/rowjay007/event-bookie/config"
)

type PaystackClient struct {
    client *resty.Client
}

type PaymentResponse struct {
    Status  bool   `json:"status"`
    Message string `json:"message"`
    Data    struct {
        AuthorizationURL string `json:"authorization_url"`
        AccessCode       string `json:"access_code"`
        Reference        string `json:"reference"`
    } `json:"data"`
}

type PaymentVerificationResponse struct {
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

func (p *PaystackClient) InitializePayment(amount int64, email string) (*PaymentResponse, error) {
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

    var paymentResponse PaymentResponse
    err = json.Unmarshal(resp.Body(), &paymentResponse)
    if err != nil {
        return nil, err
    }

    return &paymentResponse, nil
}

func (p *PaystackClient) VerifyPayment(reference string) (*PaymentVerificationResponse, error) {
    resp, err := p.client.R().
        Get(fmt.Sprintf("/transaction/verify/%s", reference))

    if err != nil {
        return nil, err
    }

    var verificationResponse PaymentVerificationResponse
    err = json.Unmarshal(resp.Body(), &verificationResponse)
    if err != nil {
        return nil, err
    }

    return &verificationResponse, nil
}
