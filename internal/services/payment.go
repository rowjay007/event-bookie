package service

import (
    "github.com/rowjay007/event-bookie/internal/models"
    "github.com/rowjay007/event-bookie/internal/repository"
)

type PaymentService struct {
    paymentRepo repository.PaymentRepository
}

func NewPaymentService(paymentRepo repository.PaymentRepository) *PaymentService {
    return &PaymentService{paymentRepo}
}

func (s *PaymentService) GetAllPayments() ([]models.Payment, error) {
    return s.paymentRepo.GetAllPayments()
}

func (s *PaymentService) GetPaymentByID(id uint) (*models.Payment, error) {
    return s.paymentRepo.GetPaymentByID(id)
}

// Implement other methods like CreatePayment, UpdatePayment, DeletePayment as needed
