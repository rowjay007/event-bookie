package models


type Organizer struct {
    ID     string  `json:"id"`
    Name   string  `json:"name"`
    Events []Event `json:"events"`
}