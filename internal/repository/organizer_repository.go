package repository

import (
    "context"
    "database/sql"
    "time"

    "github.com/rowjay007/event-bookie/internal/models"
)

type PaymentRepository struct {
    db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
    return &PaymentRepository{db}
}

func (r *PaymentRepository) CreatePayment(ctx context.Context, payment *models.Payment) error {
    query := `
        INSERT INTO payments (user_id, amount, payment_date, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
    `
    _, err := r.db.ExecContext(ctx, query, payment.UserID, payment.Amount, payment.PaymentDate, time.Now(), time.Now())
    if err != nil {
        return err
    }
    return nil
}

// Implement other methods like GetPaymentsByUserID, UpdatePayment, DeletePayment, etc.

