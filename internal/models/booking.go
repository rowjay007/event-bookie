package models

type Booking struct {
    ID        uint    `json:"id" gorm:"primaryKey"`
    EventID   uint    `json:"event_id"`
    UserID    uint    `json:"user_id"`
    PaymentID uint    `json:"payment_id"`
    Event     Event   `json:"event" gorm:"foreignKey:EventID"`
    User      User    `json:"user" gorm:"foreignKey:UserID"`
    Payment   Payment `json:"payment" gorm:"foreignKey:PaymentID"`
}
