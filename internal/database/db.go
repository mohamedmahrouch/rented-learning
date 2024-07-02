package database

import (
    "fmt"
    "os"

    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq" // Pilote PostgreSQL
)

var (
    db *sqlx.DB
)


func InitDB() (*sqlx.DB, error) {
    dbURL := os.Getenv("DATABASE_URL") // Récupérer l'URL de la base de données depuis les variables d'environnement

    // Ouvrir une connexion à la base de données
    conn, err := sqlx.Open("postgres", dbURL)
    if err != nil {
        return nil, fmt.Errorf("erreur lors de l'ouverture de la connexion à la base de données : %v", err)
    }

    // Vérifier la connexion à la base de données
    err = conn.Ping()
    if err != nil {
        return nil, fmt.Errorf("erreur lors de la vérification de la connexion à la base de données : %v", err)
    }

    db = conn
    return db, nil
}

// GetDB retourne l'instance de la connexion à la base de données
func GetDB() *sqlx.DB {
    return db
}
