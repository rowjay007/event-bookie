// payment.go
package models

// Payment represents a payment entity
type Payment struct {
    ID        string  `json:"id" db:"id,primarykey"`
    BookingID string  `json:"booking_id" db:"booking_id"`
    Amount    float64 `json:"amount" db:"amount"`
}
