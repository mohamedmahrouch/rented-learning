package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwtKey représente la clé secrète utilisée pour signer les JWT
var jwtKey = []byte("your_secret_key")

// Claims représente la structure des revendications JWT personnalisées
type Claims struct {
	UserID             int `json:"user_id"` // ID de l'utilisateur stocké dans le JWT
	jwt.StandardClaims     // Revendications standard JWT
}

// GenerateJWT génère un nouveau JWT pour l'utilisateur avec une expiration de 24 heures
func GenerateJWT(userID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Crée un nouveau token JWT signé avec la méthode de signature HS256 et les revendications définies
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseJWT vérifie et parse un token JWT donné et retourne les revendications (claims)
func ParseJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}

	// Parse le token JWT avec les revendications dans la structure Claims, en utilisant jwtKey pour la vérification de la signature
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Vérifie si le token est valide
	if !token.Valid {
		return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorExpired)
	}

	return claims, nil
}
