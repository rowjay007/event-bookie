// venue.go
package models

// Venue represents a venue entity
type Venue struct {
    ID       string `json:"id" db:"id,primarykey"`
    Name     string `json:"name" db:"name"`
    Location string `json:"location" db:"location"`
}
