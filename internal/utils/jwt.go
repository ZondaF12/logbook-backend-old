package utils

import (
	"os"
	"time"

	"github.com/ZondaF12/logbook-backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(user models.User) (string, error) {
	// Set custom claims
	claims := &models.JwtCustomClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 3)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	jwt, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return jwt, nil
}
