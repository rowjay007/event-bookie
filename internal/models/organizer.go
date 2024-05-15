// organizer.go
package models

// Organizer represents an organizer entity
type Organizer struct {
    ID   string `json:"id" db:"id,primarykey"`
    Name string `json:"name" db:"name"`
}
