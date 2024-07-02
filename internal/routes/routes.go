package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/mohamedmahrouch/rented/internal/handlers"
)

// SetupRoutes configure toutes les routes pour l'application
func SetupRoutes(r *gin.Engine, db *sqlx.DB) {
	// Routes publiques
	r.POST("/register", handlers.Register) // Endpoint pour l'inscription d'un nouvel utilisateur
	r.POST("/login", handlers.Login)       // Endpoint pour la connexion d'un utilisateur

	// Routes nécessitant une authentification
	auth := r.Group("/")
	auth.Use() // Middleware d'authentification à ajouter ici si nécessaire
	{
		auth.GET("/profile", handlers.GetUserProfile)        // Endpoint pour récupérer le profil de l'utilisateur
		auth.POST("/houses", handlers.CreateHouse)           // Endpoint pour créer une nouvelle maison
		auth.GET("/houses", handlers.GetHouses)              // Endpoint pour récupérer toutes les maisons
		auth.PUT("/houses/:id", handlers.UpdateHouse)        // Endpoint pour mettre à jour une maison spécifique
		auth.DELETE("/houses/:id", handlers.DeleteHouse)     // Endpoint pour supprimer une maison spécifique
		auth.GET("/owner/houses", handlers.GetHousesByOwner) // Endpoint pour récupérer les maisons d'un propriétaire spécifique
	}
}
