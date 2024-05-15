// user.go
package models

// User represents a user entity
type User struct {
    ID       string `json:"id" db:"id,primarykey"`
    Name     string `json:"name" db:"name"`
    Email    string `json:"email" db:"email"`
    Password string `json:"password" db:"password"`
}
