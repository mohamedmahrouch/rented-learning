// internal/models/house.go
package models

// House represents a house listing
type House struct {
	ID          int    `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Price        float64    `db:"price" json:"price"`
	OwnerID     int    `db:"owner_id" json:"owner_id"`
}
