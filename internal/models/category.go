// category.go
package models

// Category represents a category entity
type Category struct {
    ID     string  `json:"id" db:"id,primarykey"`
    Name   string  `json:"name" db:"name"`
}
