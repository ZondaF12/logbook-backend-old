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
//	@Tags			user
//	@Success		200 {object} models.SelfUser
//	@Router			/auth/self [get]
func (s *Server) GetSelfHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)

	res := s.db.GetUserByID(claims.ID)

	session := models.SelfUser{
		ID:    res.ID,
		Email: res.Email,
		Name:  res.Name,
		Role:  res.Role,
	}

	return c.JSON(http.StatusOK, session)
}

// Users godoc
//
//	@Summary		Get All Users
//	@Description	Returns a list of all users
//	@Tags			user
//	@Success		200 {array} models.SelfUser
//	@Router			/auth/users [get]
func (s *Server) GetUsersHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)

	if claims.Role != "admin" {
		return c.JSON(http.StatusForbidden, map[string]string{
			"message": "You are not authorised to view this resource",
		})
	}

	users := s.db.GetUsers()
	return c.JSON(http.StatusOK, users)
}
