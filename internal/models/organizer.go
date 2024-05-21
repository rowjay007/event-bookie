package models

import (
	"time"
)

// Organizer represents an organizer entity
type Organizer struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Contact   string    `json:"contact"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
