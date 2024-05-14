package models

import (
    "time"
)

type Event struct {
    ID          uint       `json:"id"`
    Name        string     `json:"name"`
    Description string     `json:"description"`
    Date        time.Time  `json:"date"`
    Location    string     `json:"location"`
    OrganizerID string     `json:"organizer_id"`
    Organizer   Organizer  `json:"organizer" db:"organizer_id" fk:"ID"`
    VenueID     string     `json:"venue_id"`
    Venue       Venue      `json:"venue" db:"venue_id" fk:"ID"`
    CategoryID  string     `json:"category_id"`
    Category    Category   `json:"category" db:"category_id" fk:"ID"`
    Bookings    []Booking  `json:"bookings"`
}