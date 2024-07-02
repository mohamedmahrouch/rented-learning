package main

import (
	"log"

	"github.com/mohamedmahrouch/rented/internal/config"
	"github.com/mohamedmahrouch/rented/internal/routes"
    

	"github.com/gin-gonic/gin"
)

func main() {
	// Création d'une nouvelle instance de moteur Gin avec les paramètres par défaut
	r := gin.Default()

	// Chargement de la configuration de l'application
	config.LoadConfig()

	// Connexion à la base de données PostgreSQL
	db, err := config.ConnectDB()
	if err != nil {
		// En cas d'erreur lors de la connexion à la base de données, afficher une erreur et quitter
		log.Fatalf("Erreur de connexion à la base de données : %v", err)
	}

	// Configuration des routes de l'API en utilisant le moteur Gin et la connexion à la base de données
	routes.SetupRoutes(r, db)

	// Démarrage du serveur HTTP sur le port 8080
	log.Fatal(r.Run(":8080"))
}
