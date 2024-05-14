package models

type Payment struct {
    ID        string  `json:"id"`
    BookingID string  `json:"booking_id"`
    Amount    float64 `json:"amount"`
}
