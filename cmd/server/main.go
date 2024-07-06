
package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mohamedmahrouch/rented/internal/config"
	"github.com/mohamedmahrouch/rented/internal/database"
	"github.com/mohamedmahrouch/rented/internal/routes"
)

func main() {

	conf := config.LoadConfig()

	db, err := database.InitDB(conf.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	router := gin.Default()

	// Apply CORS middleware to allow all origins
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept", "Authorization", "Access-Control-Allow-Origin"},
	}))

	routes.SetupRoutes(router, db)

	err = router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
