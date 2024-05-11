// package models

// import (
//     "time"
// )

// type Event struct {
//     ID          uint       `json:"id" gorm:"primaryKey"`
//     Name        string     `json:"name"`
//     Description string     `json:"description"`
//     Date        time.Time  `json:"date"`
//     Location    string     `json:"location"`
//     OrganizerID uint       `json:"organizer_id"`
//     Organizer   Organizer  `json:"organizer" gorm:"foreignKey:OrganizerID"`
//     VenueID     uint       `json:"venue_id"`
//     Venue       Venue      `json:"venue" gorm:"foreignKey:VenueID"`
//     CategoryID  uint       `json:"category_id"`
//     Category    Category   `json:"category" gorm:"foreignKey:CategoryID"`
//     Bookings    []Booking  `json:"bookings"`
// }
