package models

type Booking struct {
    ID        string  `json:"id"`
    EventID   string  `json:"event_id"`
    UserID    string  `json:"user_id"`
    PaymentID string  `json:"payment_id"`
    Event     Event   `json:"event" db:"event_id" fk:"ID"`
    User      User    `json:"user" db:"user_id" fk:"ID"`
    Payment   Payment `json:"payment" db:"payment_id" fk:"ID"`
}

