package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	BookingID  uint           `gorm:"not null" json:"booking_id"`
	Amount     float64        `gorm:"not null" json:"amount"`
	PaidAt     time.Time      `gorm:"not null" json:"paid_at"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	Booking    Booking        `gorm:"foreignKey:BookingID" json:"booking"`
}
