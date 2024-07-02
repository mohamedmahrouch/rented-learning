package handlers

import (
	"net/http"
	"strconv"

	"github.com/mohamedmahrouch/rented/internal/models"

	"github.com/gin-gonic/gin"
)

// GetUserProfile récupère le profil de l'utilisateur
func GetUserProfile(c *gin.Context) {
	// Récupère l'ID de l'utilisateur à partir du contexte
	userID, _ := strconv.Atoi(c.GetString("userID"))

	// Récupère l'utilisateur depuis la base de données en utilisant son ID
	user, err := models.GetUserByID(userID)
	if err != nil {
		// Si une erreur se produit lors de la récupération de l'utilisateur, retourne une réponse HTTP avec le statut InternalServerError
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Réponse réussie avec les détails de l'utilisateur et le statut OK si tout se passe bien
	c.JSON(http.StatusOK, gin.H{"user": user})
}
