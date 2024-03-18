package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Session struct {
	UserID      uuid.UUID `json:"user_id"`
	Token       uuid.UUID `json:"access_token"`
	ExpiryDate  time.Time `json:"expiry_date"`
	CreatedAt   time.Time `json:"created_at"`
	RefreshDate time.Time `json:"refresh_date"`
}

type JwtCustomClaims struct {
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}
