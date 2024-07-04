// internal/handlers/user.go
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/mohamedmahrouch/rented/internal/models"
)

// UserHandler handles user related requests
type UserHandler struct {
	DB *sqlx.DB
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(db *sqlx.DB) *UserHandler {
	return &UserHandler{DB: db}
}

// GetUserByID retrieves a user by their ID
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	err := h.DB.Get(&user, "SELECT id, name, email FROM users WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser updates a user's information
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID, _ = strconv.Atoi(id)

	_, err := h.DB.NamedExec(`UPDATE users SET name=:name, email=:email WHERE id=:id`, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser deletes a user
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	_, err := h.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
