package server

import (
	"net/http"

	"github.com/ZondaF12/logbook-backend/internal/database"
	"github.com/labstack/echo/v4"
)

type Self struct {
	SelfUser
	database.Session
}

type SelfUser struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}


// Self godoc
//
//	@Summary		Get Authenticated User
//	@Description	Returns the authenticated user
//	@Tags			auth
//	@Success		200
//	@Router			/auth/self [get]
func (s *Server) GetSelfHandler(c echo.Context) error {
	email, _, _ := c.Request().BasicAuth()

	res := s.db.GetUserByEmail(email)

	user := SelfUser{
		Email: res.Email,
		Name:  res.Name,
	}

	resp := Self{
		SelfUser: user,
		Session:  c.Get("session").(database.Session),
	}

	return c.JSON(http.StatusOK, resp)
}
