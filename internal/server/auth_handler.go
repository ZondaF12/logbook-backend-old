package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ZondaF12/logbook-backend/internal/auth"
	"github.com/ZondaF12/logbook-backend/internal/database"
	"github.com/labstack/echo/v4"
)

// Register godoc
//
//	@Summary		Register Route
//	@Description	Register a new user
//	@Tags			auth
//	@Success		200
//	@Router			/auth/register [post]
func (s *Server) RegisterHandler(c echo.Context) error {
	newUser := database.User{} // Slice of User instances

	err := json.NewDecoder(c.Request().Body).Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		return err
	}

	res := s.db.AddUserToDB(newUser)
	if res["message"] == "User already exists" {
		return c.JSON(http.StatusForbidden, res)
	}

	return c.JSON(http.StatusOK, res)
}

// Login godoc
//
//	@Summary		Login Route
//	@Description	Login a user
//	@Tags			auth
//	@Success		200
//	@Router			/auth/login [post]
func (s *Server) LoginHandler(c echo.Context) error {
	user := database.User{} // Slice of User instances

	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		return err
	}

	res := s.db.GetUserByEmail(user.Email)
	valid := auth.ComparePasswords(res.Password, []byte(user.Password))

	if !valid {
		return c.JSON(http.StatusForbidden, map[string]string{
			"message": "Invalid Credentials",
		})
	}

	session := s.db.CreateSession(res)

	return c.JSON(http.StatusOK, session)
}