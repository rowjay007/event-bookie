package models

type Organizer struct {
    ID     uint    `json:"id" gorm:"primaryKey"`
    Name   string  `json:"name"`
    Events []Event `json:"events"`
}
