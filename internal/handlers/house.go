package handlers

import (
	"net/http"
	"strconv"

	"github.com/mohamedmahrouch/rented/internal/models"

	"github.com/gin-gonic/gin"
)

// CreateHouse crée une nouvelle maison
func CreateHouse(c *gin.Context) {
	var house models.House
	// Bind les données JSON envoyées par le client à la structure models.House
	if err := c.ShouldBindJSON(&house); err != nil {
		// Si une erreur se produit lors du binding, retourne une réponse HTTP avec le statut BadRequest
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Récupère l'ID du propriétaire à partir du contexte et l'assigne à la maison
	ownerID, _ := strconv.Atoi(c.GetString("userID"))
	house.OwnerID = ownerID

	// Création de la maison dans la base de données
	_, err := models.CreateHouse(house)
	if err != nil {
		// Si une erreur se produit lors de la création de la maison, retourne une réponse HTTP avec le statut InternalServerError
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Réponse réussie avec le statut Created si tout se passe bien
	c.JSON(http.StatusCreated, gin.H{"message": "House created successfully"})
}

// GetHouses récupère toutes les maisons
func GetHouses(c *gin.Context) {
	// Récupère toutes les maisons depuis la base de données
	houses, err := models.GetAllHouses()
	if err != nil {
		// Si une erreur se produit lors de la récupération des maisons, retourne une réponse HTTP avec le statut InternalServerError
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Réponse réussie avec la liste des maisons et le statut OK si tout se passe bien
	c.JSON(http.StatusOK, gin.H{"houses": houses})
}

// UpdateHouse met à jour une maison existante
func UpdateHouse(c *gin.Context) {
	// Récupère l'ID de la maison à mettre à jour depuis les paramètres de l'URL
	id, _ := strconv.Atoi(c.Param("id"))

	var house models.House
	// Bind les données JSON envoyées par le client à la structure models.House
	if err := c.ShouldBindJSON(&house); err != nil {
		// Si une erreur se produit lors du binding, retourne une réponse HTTP avec le statut BadRequest
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Met à jour la maison dans la base de données
	err := models.UpdateHouse(id, house)
	if err != nil {
		// Si une erreur se produit lors de la mise à jour de la maison, retourne une réponse HTTP avec le statut InternalServerError
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Réponse réussie avec un message et le statut OK si tout se passe bien
	c.JSON(http.StatusOK, gin.H{"message": "House updated successfully"})
}

// DeleteHouse supprime une maison existante
func DeleteHouse(c *gin.Context) {
	// Récupère l'ID de la maison à supprimer depuis les paramètres de l'URL
	id, _ := strconv.Atoi(c.Param("id"))

	// Supprime la maison de la base de données
	err := models.DeleteHouse(id)
	if err != nil {
		// Si une erreur se produit lors de la suppression de la maison, retourne une réponse HTTP avec le statut InternalServerError
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Réponse réussie avec un message et le statut OK si tout se passe bien
	c.JSON(http.StatusOK, gin.H{"message": "House deleted successfully"})
}

// GetHousesByOwner récupère toutes les maisons appartenant à un propriétaire spécifique
func GetHousesByOwner(c *gin.Context) {
	// Récupère l'ID du propriétaire à partir du contexte
	ownerID, _ := strconv.Atoi(c.GetString("userID"))

	// Récupère toutes les maisons appartenant à ce propriétaire depuis la base de données
	houses, err := models.GetHousesByOwner(ownerID)
	if err != nil {
		// Si une erreur se produit lors de la récupération des maisons, retourne une réponse HTTP avec le statut InternalServerError
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Réponse réussie avec la liste des maisons et le statut OK si tout se passe bien
	c.JSON(http.StatusOK, gin.H{"houses": houses})
}
