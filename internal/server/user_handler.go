package server

import (
	"net/http"

	"github.com/ZondaF12/logbook-backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// Self godoc
//
//	@Summary		Get Authenticated User
//	@Description	Returns the authenticated user
//	@Tags			auth
//	@Success		200
//	@Router			/auth/self [get]
func (s *Server) GetSelfHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)

	res := s.db.GetUserByID(claims.ID)

	session := models.SelfUser{
		ID:    res.ID,
		Email: res.Email,
		Name:  res.Name,
	}

	return c.JSON(http.StatusOK, session)
}
