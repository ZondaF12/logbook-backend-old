package server

import (
	"net/http"

	"github.com/ZondaF12/logbook-backend/internal/models"
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
	email, _, _ := c.Request().BasicAuth()

	res := s.db.GetUserByEmail(email)

	user := models.SelfUser{
		Email: res.Email,
		Name:  res.Name,
	}

	resp := models.Self{
		SelfUser: user,
		Session:  c.Get("session").(models.Session),
	}

	return c.JSON(http.StatusOK, resp)
}
