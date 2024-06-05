package payment

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PaystackService struct {
	secretKey string
	client    *http.Client
}

func NewPaystackService(secretKey string) *PaystackService {
	return &PaystackService{
		secretKey: secretKey,
		client:    &http.Client{Timeout: 10 * time.Second},
	}
}

func (p *PaystackService) InitializePayment(ctx context.Context, amount float64, email string, reference string) (string, error) {
	url := "https://api.paystack.co/transaction/initialize"
	payload := map[string]interface{}{
		"reference": reference,
		"amount":    int(amount * 100), 
		"email":     email,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error marshaling request payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+p.secretKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
		Data    struct {
			AuthorizationURL string `json:"authorization_url"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("error decoding response: %w", err)
	}

	if !response.Status {
		return "", fmt.Errorf("error initializing payment: %s", response.Message)
	}

	return response.Data.AuthorizationURL, nil
}

// VerifyPayment verifies the status of a payment with the Paystack payment gateway.
func (p *PaystackService) VerifyPayment(ctx context.Context, reference string) error {
	url := fmt.Sprintf("https://api.paystack.co/transaction/verify/%s", reference)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+p.secretKey)

	resp, err := p.client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
		Data    struct {
			Status          string `json:"status"`
			GatewayResponse string `json:"gateway_response"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	if !response.Status {
		return fmt.Errorf("error verifying payment: %s", response.Message)
	}

	if response.Data.Status != "success" {
		return fmt.Errorf("payment verification failed: %s", response.Data.GatewayResponse)
	}

	return nil
}
