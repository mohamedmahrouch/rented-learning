// internal/routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/mohamedmahrouch/rented/internal/handlers"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(router *gin.Engine, db *sqlx.DB) {
	authHandler := handlers.NewAuthHandler(db)
	houseHandler := handlers.NewHouseHandler(db)
	userHandler := handlers.NewUserHandler(db)

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)
    

	router.GET("/houses", houseHandler.GetHouses)
	router.POST("/houses", houseHandler.CreateHouse)
	router.GET("/houses/:id", houseHandler.GetHouseByID)
	router.PUT("/houses/:id", houseHandler.UpdateHouse)
	router.DELETE("/houses/:id", houseHandler.DeleteHouse)
	router.GET("/houses/user/:id", houseHandler.GetHousesByOwner)

	router.GET("/users/:id", userHandler.GetUserByID)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)
}
