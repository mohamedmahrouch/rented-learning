package handlers

import (
    "net/http"
    "github.com/mohamedmahrouch/rented/internal/models"
    "github.com/mohamedmahrouch/rented/internal/utils"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

// Register permet d'enregistrer un nouvel utilisateur
func Register(c *gin.Context) {
    var user models.User
    // Bind les données JSON envoyées par le client à la structure models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        // Si une erreur se produit lors du binding, retourne une réponse HTTP avec le statut BadRequest
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash du mot de passe utilisateur avant de l'enregistrer dans la base de données
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        // En cas d'erreur lors du hashage du mot de passe, retourne une réponse HTTP avec le statut InternalServerError
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
        return
    }
    user.Password = string(hashedPassword)

    // Création de l'utilisateur dans la base de données
    _, err = models.CreateUser(user)
    if err != nil {
        // Si une erreur se produit lors de la création de l'utilisateur, retourne une réponse HTTP avec le statut InternalServerError
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Réponse réussie avec le statut Created si tout se passe bien
    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login permet à un utilisateur de se connecter
func Login(c *gin.Context) {
    var input models.User
    // Bind les données JSON envoyées par le client à la structure models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        // Si une erreur se produit lors du binding, retourne une réponse HTTP avec le statut BadRequest
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Récupération de l'utilisateur depuis la base de données en utilisant son adresse email
    user, err := models.GetUserByEmail(input.Email)
    if err != nil {
        // Si aucun utilisateur correspondant n'est trouvé, retourne une réponse HTTP avec le statut Unauthorized
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    // Vérification du mot de passe en comparant le hash stocké avec le mot de passe fourni par l'utilisateur
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
    if err != nil {
        // Si le mot de passe ne correspond pas, retourne une réponse HTTP avec le statut Unauthorized
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    // Génération d'un token JWT pour l'utilisateur authentifié
    token, err := utils.GenerateJWT(user.ID)
    if err != nil {
        // En cas d'erreur lors de la génération du token JWT, retourne une réponse HTTP avec le statut InternalServerError
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    // Réponse réussie avec le token JWT si tout se passe bien
    c.JSON(http.StatusOK, gin.H{"token": token})
}
