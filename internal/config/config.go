package config

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	// Importation du pilote PostgreSQL
	_ "github.com/lib/pq"
)

// LoadConfig charge les variables d'environnement à partir d'un fichier .env
func LoadConfig() {
	// Chargement du fichier .env
	err := godotenv.Load()
	if err != nil {
		// Affiche un message d'erreur si le fichier .env ne peut pas être chargé
		log.Println("Error loading .env file")
	}
}

// ConnectDB établit une connexion à la base de données PostgreSQL
func ConnectDB() (*sqlx.DB, error) {
	// Connexion à la base de données en utilisant la variable d'environnement DATABASE_URL
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		// Retourne une erreur si la connexion échoue
		return nil, err
	}
	// Retourne l'objet de connexion à la base de données en cas de succès
	return db, nil
}
