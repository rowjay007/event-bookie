package models

type Venue struct {
    ID       uint    `json:"id" gorm:"primaryKey"`
    Name     string  `json:"name"`
    Location string  `json:"location"`
    Events   []Event `json:"events"`
}
