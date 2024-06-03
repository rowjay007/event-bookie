package repository

import (
    "github.com/rowjay007/event-bookie/internal/models"
    "gorm.io/gorm"
)

type PaymentRepository struct {
    DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
    return &PaymentRepository{DB: db}
}

func (pr *PaymentRepository) Create(payment *models.Payment) error {
    return pr.DB.Create(payment).Error
}

func (pr *PaymentRepository) GetByID(id uint) (*models.Payment, error) {
    var payment models.Payment
    if err := pr.DB.First(&payment, id).Error; err != nil {
        return nil, err
    }
    return &payment, nil
}

func (pr *PaymentRepository) Update(payment *models.Payment) error {
    return pr.DB.Save(payment).Error
}



func (pr *PaymentRepository) GetAll(queryParams map[string]string, offset, limit int) ([]models.Payment, int64, error) {
    var payments []models.Payment
    var total int64

    query := pr.DB.Model(&models.Payment{})

    // Implement any filtering logic here

    if err := query.Count(&total).Error; err != nil {
        return nil, 0, err
    }

    if offset >= 0 && limit > 0 {
        query = query.Offset(offset).Limit(limit)
    }

    if err := query.Find(&payments).Error; err != nil {
        return nil, 0, err
    }

    return payments, total, nil
}

func (pr *PaymentRepository) Delete(id uint) error {
    return pr.DB.Delete(&models.Payment{}, id).Error
}