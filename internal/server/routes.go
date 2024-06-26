package server

import (
	"net/http"
	"os"

	_ "github.com/ZondaF12/logbook-backend/docs"
	"github.com/ZondaF12/logbook-backend/internal/models"
	"github.com/golang-jwt/jwt/v5"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

// @title			Logbook Backend API
// @version		1.0
// @description	Backend for the mobile app `Logbook` made with Echo and Golang
// @host			localhost:8080
// @BasePath		/
func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	// Default Routes
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", s.HelloWorldHandler)
	e.GET("/health", s.HealthHandler)
	e.POST("/register", s.RegisterHandler)

	login := e.Group("/login")
	login.Use(middleware.BasicAuth(s.UserAuthenticateByCredentials))
	login.POST("", s.LoginHandler)

	// Auth Routes
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}
	auth := e.Group("/auth")
	auth.Use(echojwt.WithConfig(config))

	// Self Routes
	auth.GET("/self", s.GetSelfHandler)
	auth.POST("/self", s.AddSelfHandler)
	auth.PUT("/self", s.UpdateSelfHandler)

	// Users Routes
	auth.GET("/users", s.GetUsersHandler)
	auth.GET("/users/:id", s.GetUserByIDHandler)

	// Vehicles Routes
	auth.GET("/vehicles", s.GetAuthenticatedUserVehiclesHandler)
	auth.POST("/vehicles", s.AddVehicleToAuthenticatedUserHandler)
	auth.GET("/vehicles/:id", s.GetVehicleByIDHandler)
	auth.GET("/vehicles/registration/:reg", s.GetAuthenticatedUsersVehicleByRegistrationHandler)

	// Utils Routes
	auth.POST("/utils/username", s.IsUsernameAvailableHandler)
	return e
}

// HelloWorld godoc
//
// @Summary		Hello World Route
// @Description	returns `Hello World`
// @Tags		default
// @Success		200 {string} message
// @Router		/ [get]
func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

// Health godoc
//
// @Summary		Returns the database health
// @Description	get the database health
// @Tags		default
// @Success		200 {string} message
// @Router		/health [get]
func (s *Server) HealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
