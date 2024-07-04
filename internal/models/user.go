// internal/models/user.go
package models

// User represents a user in the system
type User struct {
	ID       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
}

// Credentials represents login credentials
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
