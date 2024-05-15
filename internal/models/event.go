// event.go
package models

import "time"

// Event represents an event entity
type Event struct {
    ID          string    `json:"id" db:"id,primarykey"`
    Name        string    `json:"name" db:"name"`
    Description string    `json:"description" db:"description"`
    Date        time.Time `json:"date" db:"date"`
    Location    string    `json:"location" db:"location"`
    OrganizerID string    `json:"organizer_id" db:"organizer_id"`
    VenueID     string    `json:"venue_id" db:"venue_id"`
    CategoryID  string    `json:"category_id" db:"category_id"`
}
