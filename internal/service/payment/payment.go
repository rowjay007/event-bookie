package service

import (
	"github.com/rowjay007/event-bookie/internal/models"
	"github.com/rowjay007/event-bookie/internal/repository"
)

type PaymentService struct {
    PaymentRepo *repository.PaymentRepository
}

func NewPaymentService(paymentRepo *repository.PaymentRepository) *PaymentService {
    return &PaymentService{PaymentRepo: paymentRepo}
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


