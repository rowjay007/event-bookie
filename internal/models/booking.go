// booking.go
package models

// Booking represents a booking entity
type Booking struct {
    ID        string `json:"id" db:"id,primarykey"`
    EventID   string `json:"event_id" db:"event_id"`
    UserID    string `json:"user_id" db:"user_id"`
    PaymentID string `json:"payment_id" db:"payment_id"`
}
