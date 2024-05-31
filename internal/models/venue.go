package models

import (
	"time"
)

type Venue struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Capacity  uint      `json:"capacity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
