package models

import (
	"time"
)

type Event struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time" time_format:"02-01-2006 15:04:05"`
	EndTime     time.Time `json:"end_time" time_format:"02-01-2006 15:04:05"`
	VenueID     uint      `json:"venue_id"`
	OrganizerID uint      `json:"organizer_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
