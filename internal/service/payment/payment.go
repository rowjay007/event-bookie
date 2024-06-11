package payment

import (
	"github.com/rowjay007/event-bookie/internal/models"
	"github.com/rowjay007/event-bookie/internal/repository"
)

type PaymentService struct {
	PaymentRepo    *repository.PaymentRepository
	PaystackClient *PaystackClient 
	FlutterwaveClient *FlutterwaveClient
}

func NewPaymentService(paymentRepo *repository.PaymentRepository, paystackClient *PaystackClient, FlutterwaveClient *FlutterwaveClient ) *PaymentService {
	return &PaymentService{
		PaymentRepo:    paymentRepo,
		PaystackClient: paystackClient, 
		FlutterwaveClient: FlutterwaveClient,

	}
}
func (ps *PaymentService) CreatePayment(payment *models.Payment) error {
    if err := ps.PaymentRepo.Create(payment); err != nil {
        return err
    }
    return nil
}

func (ps *PaymentService) GetAllPayments(params map[string]string, offset, limit int, sort, order, filter string) ([]models.Payment, int64, error) {
    return ps.PaymentRepo.GetAll(params, offset, limit, sort, order, filter)
}

func (ps *PaymentService) GetPaymentByID(id uint) (*models.Payment, error) {
    return ps.PaymentRepo.GetByID(id)
}

func (ps *PaymentService) UpdatePayment(payment *models.Payment) error {
    return ps.PaymentRepo.Update(payment)
}

func (ps *PaymentService) DeletePayment(id uint) error {
    return ps.PaymentRepo.Delete(id)
}
func (ps *PaymentService) InitiatePaystackPayment(amount int64, email string) (*PaystackPaymentResponse, error) {
	return ps.PaystackClient.InitializePaystackPayment(amount, email)
}

func (ps *PaymentService) VerifyPaystackPayment(reference string) (*PaystackVerificationResponse, error) {
	return ps.PaystackClient.VerifyPaystackPayment(reference)
}

func (ps *PaymentService) InitiateFlutterwavePayment(amount int64, email, txRef string) (*FlutterwavePaymentResponse, error) {
	return ps.FlutterwaveClient.InitializePayment(amount, email, txRef)
}

func (ps *PaymentService) VerifyFlutterwavePayment(txRef string) (*FlutterwavePaymentVerificationResponse, error) {
	return ps.FlutterwaveClient.VerifyPayment(txRef)
}
