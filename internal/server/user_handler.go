package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ZondaF12/logbook-backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *Server) AddNewUserHandler(c echo.Context) error {
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
//	@Tags			user
//	@Success		200 {object} models.User
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

// Users godoc
//
//	@Summary		Get All Users
//	@Description	Returns a list of all users
//	@Tags			user
//	@Success		200 {array} models.User
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

// Users godoc
//
//	@Summary		Get a User by ID
//	@Description	Returns a user object
//	@Tags			user
//	@Success		200 {object} models.User
//	@Error			403 {string} message
//	@Router			/auth/users/:id [get]
func (s *Server) GetUserByIDHandler(c echo.Context) error {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	user, err := s.db.GetUserByID(uuid)
	if err != nil {
		return c.JSON(http.StatusForbidden, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (s *Server) UpdateUserByIDHandler(c echo.Context) error {
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
