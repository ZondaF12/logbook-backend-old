package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ZondaF12/logbook-backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// Add Self godoc
//
//	@Summary		Add New Authenticated User
//	@Description	Creates a new user
//	@Tags			self
//	@Param			request body models.UpdateSelf true "update params"
//	@Success		200 {object} models.User
//	@Failure		403 {string} message
//	@Router			/auth/self [post]
func (s *Server) AddSelfHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)

	newUser := models.User{
		ID: claims.ID,
	}

	err := json.NewDecoder(c.Request().Body).Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		return err
	}

	res := s.db.AddNewUserToDB(newUser)

	return c.JSON(http.StatusOK, res)
}

// Self godoc
//
//	@Summary		Get Authenticated User
//	@Description	Returns the authenticated user
//	@Tags			self
//	@Success		200 {object} models.User
//	@Failure		403 {string} message
//	@Router			/auth/self [get]
func (s *Server) GetSelfHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)

	self, err := s.db.GetUserByID(claims.ID)
	if err != nil {
		return c.JSON(http.StatusForbidden, err.Error())
	}

	return c.JSON(http.StatusOK, self)
}

// Update Self godoc
//
//	@Summary		Updates the Authenticated User
//	@Description	updates the authenticated user
//	@Tags			self
//	@Success		200 {object} models.User
//	@Failure		403 {string} message
//	@Router			/auth/self [put]
func (s *Server) UpdateSelfHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)

	newSelf := models.User{
		ID: claims.ID,
	}

	err := json.NewDecoder(c.Request().Body).Decode(&newSelf)
	if err != nil {
		fmt.Println(err)
		return err
	}

	updatedSelf, err := s.db.UpdateUserByID(claims.ID, newSelf)
	if err != nil {
		return c.JSON(http.StatusForbidden, err.Error())
	}

	return c.JSON(http.StatusOK, updatedSelf)
}

// Valid Username godoc
//
//	@Summary		Is a Username Available
//	@Description	Checks if a username is available
//	@Tags			utils
//	@Success		200 {string} bool
//	@Router			/auth/utils/username [get]
func (s *Server) IsUsernameAvailableHandler(c echo.Context) error {
	username := models.Username{}

	err := json.NewDecoder(c.Request().Body).Decode(&username)
	if err != nil {
		fmt.Println(err)
		return err
	}

	exists := s.db.IsUsernameAvailable(username.Username)

	return c.JSON(http.StatusOK, exists)
}
