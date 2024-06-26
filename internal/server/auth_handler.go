package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ZondaF12/logbook-backend/internal/auth"
	"github.com/ZondaF12/logbook-backend/internal/models"
	"github.com/ZondaF12/logbook-backend/internal/utils"
	"github.com/labstack/echo/v4"
)

// Register godoc
//
//	@Summary		Register Route
//	@Description	Register a new user
//	@Tags			auth
//	@Success		200 {string} message
//	@Router			/register [post]
func (s *Server) RegisterHandler(c echo.Context) error {
	newUser := models.UserAuth{} // Slice of User instances

	err := json.NewDecoder(c.Request().Body).Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		return err
	}

	res := s.db.AddAuthenticationToDB(newUser)
	if res["message"] == "User already exists" {
		return c.JSON(http.StatusForbidden, res)
	}

	return c.JSON(http.StatusOK, res)
}

// Login godoc
//
//	@Summary		Login Route
//	@Description	Logs in a user
//	@Tags			auth
//	@Success		200 {object} models.Token
//	@Router			/login [post]
func (s *Server) LoginHandler(c echo.Context) error {
	token := c.Get("token")
	return c.JSON(http.StatusOK, models.Token{
		Token: token.(string),
	})
}

func (s *Server) UserAuthenticateByCredentials(username, password string, c echo.Context) (bool, error) {
	res := s.db.GetAuthenticationByEmail(username)
	valid := auth.ComparePasswords(res.Password, []byte(password))

	if !valid {
		return false, c.JSON(http.StatusForbidden, map[string]string{
			"message": "Invalid Credentials",
		})
	}

	token, err := utils.GenerateToken(res)
	if err != nil {
		return false, c.JSON(http.StatusForbidden, map[string]string{
			"message": "Invalid Credentials",
		})
	}

	c.Set("token", token)

	return true, nil
}
