// internal/handlers/house.go
package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/mohamedmahrouch/rented/internal/models"
)

// HouseHandler handles house related requests
type HouseHandler struct {
	DB *sqlx.DB
}

// NewHouseHandler creates a new HouseHandler
func NewHouseHandler(db *sqlx.DB) *HouseHandler {
	return &HouseHandler{DB: db}
}

// GetHouses retrieves the list of houses
func (h *HouseHandler) GetHouses(c *gin.Context) {
	var houses []models.House
	err := h.DB.Select(&houses, "SELECT * FROM houses")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve houses"})
		return
	}

	c.JSON(http.StatusOK, houses)
}

// CreateHouse creates a new house listing
func (h *HouseHandler) CreateHouse(c *gin.Context) {
	var house models.House
	if err := c.ShouldBindJSON(&house); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.DB.NamedExec(`INSERT INTO houses (title, description, price, owner_id) VALUES (:title, :description, :price, :owner_id)`, &house)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create house listing"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "House listing created successfully"})
}

// GetHouseByID retrieves a house by its ID
func (h *HouseHandler) GetHouseByID(c *gin.Context) {
	id := c.Param("id")

	var house models.House
	err := h.DB.Get(&house, "SELECT * FROM houses WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve house by its ID"})
		return
	}

	
	c.JSON(http.StatusOK, house)

}

// UpdateHouse updates a house listing
func (h *HouseHandler) UpdateHouse(c *gin.Context) {
	id := c.Param("id")
	var house models.House
	if err := c.ShouldBindJSON(&house); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	house.ID, _ = strconv.Atoi(id)

	_, err := h.DB.NamedExec(`UPDATE houses SET title=:title, description=:description, price=:price WHERE id=:id`, &house)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update house listing"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "House listing updated successfully"})
}

// DeleteHouse deletes a house listing
func (h *HouseHandler) DeleteHouse(c *gin.Context) {
	id := c.Param("id")

	_, err := h.DB.Exec("DELETE FROM houses WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete house listing"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "House listing deleted successfully"})
}

// GetHousesByOwner retrieves the list of houses by owner ID
func (h *HouseHandler) GetHousesByOwner(c *gin.Context) {
	ownerID := c.Param("id")
	fmt.Println(ownerID)
	var houses []models.House
	err := h.DB.Select(&houses, "SELECT * FROM houses WHERE owner_id=$1", ownerID)
	fmt.Println(err)
	if err != nil {
		fmt.Println("hhhhhhh")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve houses  by owner ID "})
		return
	}
	response := gin.H{
		"message": "Houses retrieved successfully by owner ID",
		"data": gin.H{
			"houses": houses, // houses représente la liste des maisons récupérées
		},
	}

	c.JSON(http.StatusOK, response)

}
