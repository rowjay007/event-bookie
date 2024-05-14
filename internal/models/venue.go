package models

type Venue struct {
    ID       string  `json:"id"`
    Name     string  `json:"name"`
    Location string  `json:"location"`
    Events   []Event `json:"events"`
}
