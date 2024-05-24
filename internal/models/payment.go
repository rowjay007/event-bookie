package models

import (
	"time"
)

// Payment represents a payment entity
type Payment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	BookingID uint      `json:"booking_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}