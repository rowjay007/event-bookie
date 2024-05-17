package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	EventID   uint           `gorm:"not null" json:"event_id"`
	Quantity  int            `gorm:"not null" json:"quantity"`
	TotalCost float64        `gorm:"not null" json:"total_cost"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	User      User           `gorm:"foreignKey:UserID" json:"user"`
	Event     Event          `gorm:"foreignKey:EventID" json:"event"`
}
