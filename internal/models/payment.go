package models

type Payment struct {
    ID        uint    `json:"id" gorm:"primaryKey"`
    BookingID uint    `json:"booking_id"`
    Amount    float64 `json:"amount"`
}
