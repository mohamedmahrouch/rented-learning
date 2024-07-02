package models

import "github.com/jmoiron/sqlx"

// House représente la structure d'une maison en base de données et en JSON
type House struct {
	ID          int    `json:"id" db:"id"`                   // ID de la maison
	Title       string `json:"title" db:"title"`             // Titre de la maison
	Description string `json:"description" db:"description"` // Description de la maison
	Price       int    `json:"price" db:"price"`             // Prix de la maison
	OwnerID     int    `json:"owner_id" db:"owner_id"`       // ID du propriétaire de la maison
}

var db *sqlx.DB // Instance de connexion à la base de données

// SetDB définit l'instance de connexion à la base de données
func SetDB(database *sqlx.DB) {
	db = database
}

// CreateHouse crée une nouvelle maison dans la base de données
func CreateHouse(house House) (int, error) {
	query := "INSERT INTO houses (title, description, price, owner_id) VALUES ($1, $2, $3, $4) RETURNING id"
	var id int
	err := db.QueryRow(query, house.Title, house.Description, house.Price, house.OwnerID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetAllHouses récupère toutes les maisons depuis la base de données
func GetAllHouses() ([]House, error) {
	var houses []House
	err := db.Select(&houses, "SELECT * FROM houses")
	if err != nil {
		return nil, err
	}
	return houses, nil
}

// UpdateHouse met à jour une maison existante dans la base de données
func UpdateHouse(id int, house House) error {
	query := "UPDATE houses SET title = $1, description = $2, price = $3 WHERE id = $4"
	_, err := db.Exec(query, house.Title, house.Description, house.Price, id)
	return err
}

// DeleteHouse supprime une maison de la base de données
func DeleteHouse(id int) error {
	query := "DELETE FROM houses WHERE id = $1"
	_, err := db.Exec(query, id)
	return err
}

// GetHousesByOwner récupère toutes les maisons appartenant à un propriétaire spécifique depuis la base de données
func GetHousesByOwner(ownerID int) ([]House, error) {
	var houses []House
	err := db.Select(&houses, "SELECT * FROM houses WHERE owner_id = $1", ownerID)
	if err != nil {
		return nil, err
	}
	return houses, nil
}
