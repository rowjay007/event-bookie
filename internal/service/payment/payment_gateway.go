package payment

import "context"

type PaymentGateway interface {
    InitializePayment(ctx context.Context, amount float64, email string, reference string) (string, error)
    VerifyPayment(ctx context.Context, reference string) error
}
