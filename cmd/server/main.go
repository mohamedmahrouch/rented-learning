
package main

import (
	"log"

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

	
	routes.SetupRoutes(router, db)

	
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
