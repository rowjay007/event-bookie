package repository

import (
	"github.com/rowjay007/event-bookie/internal/models"
	"gorm.io/gorm"
)

type PaymentRepository struct {
    db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
    return &PaymentRepository{db}
}

func (r *PaymentRepository) GetAllPayments() ([]models.Payment, error) {
    var payments []models.Payment
    if err := r.db.Find(&payments).Error; err != nil {
        return nil, err
    }
    return payments, nil
}

func (r *PaymentRepository) GetPaymentByID(id uint) (*models.Payment, error) {
    var payment models.Payment
    if err := r.db.First(&payment, id).Error; err != nil {
        return nil, err
    }
    return &payment, nil
}

// Implement other methods like CreatePayment, UpdatePayment, DeletePayment as needed
