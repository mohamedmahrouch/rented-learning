
package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/mohamedmahrouch/rented/internal/models"
	"github.com/mohamedmahrouch/rented/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

// AuthHandler gère les requêtes liées à l'authentification
type AuthHandler struct {
	DB *sqlx.DB
}

// NewAuthHandler crée un nouveau AuthHandler
func NewAuthHandler(db *sqlx.DB) *AuthHandler {
	return &AuthHandler{DB: db}
}

// Register gère l'enregistrement des utilisateurs
func (h *AuthHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON invalide: " + err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Échec du hachage du mot de passe: " + err.Error()})
		return
	}
	user.Password = string(hashedPassword)
	fmt.Println("User after password hashing:", user)
	_, err = h.DB.NamedExec(`INSERT INTO users (name, email, password) VALUES (:name, :email, :password)`, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Échec de l'enregistrement de l'utilisateur dans rented: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enregistrement réussi"})
}

// Login gère la connexion des utilisateurs
func (h *AuthHandler) Login(c *gin.Context) {
	var credentials models.Credentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON invalide: " + err.Error()})
		return
	}

	var user models.User
	err := h.DB.Get(&user, "SELECT id ,password  FROM users WHERE email=$1", credentials.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email ou mot de passe incorrect"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email ou mot de passe invalide"})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Échec de la génération du jeton: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"message": "Bienvenue vos information est correct!" ,"token": token})
	
}
