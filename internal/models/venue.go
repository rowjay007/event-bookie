package models

import (
	"time"

	"gorm.io/gorm"
)

type Venue struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Address   string         `gorm:"not null" json:"address"`
	City      string         `gorm:"not null" json:"city"`
	State     string         `json:"state"`
	ZipCode   string         `json:"zip_code"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Events    []Event        `json:"events"`
}
