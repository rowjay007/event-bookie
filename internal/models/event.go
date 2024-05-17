package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	StartDate   time.Time      `gorm:"not null" json:"start_date"`
	EndDate     time.Time      `gorm:"not null" json:"end_date"`
	VenueID     uint           `gorm:"not null" json:"venue_id"`
	OrganizerID uint           `gorm:"not null" json:"organizer_id"`
	CategoryID  uint           `gorm:"not null" json:"category_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Venue       Venue          `gorm:"foreignKey:VenueID" json:"venue"`
	Organizer   Organizer      `gorm:"foreignKey:OrganizerID" json:"organizer"`
	Category    Category       `gorm:"foreignKey:CategoryID" json:"category"`
}
