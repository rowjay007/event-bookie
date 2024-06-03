package models

import (
    "time"
)

type Payment struct {
    ID            uint      `json:"id" gorm:"primaryKey"`
    UserID        uint      `json:"user_id"`
    BookingID     uint      `json:"booking_id"`
    Amount        float64   `json:"amount"`
    PaymentMethod string    `json:"payment_method"`
    Status        string    `json:"status"`
    TransactionID string    `json:"transaction_id"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
}
