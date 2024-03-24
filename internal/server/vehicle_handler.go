package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ZondaF12/logbook-backend/internal/models"
	"github.com/ZondaF12/logbook-backend/internal/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Add Vehicle godoc
//
//	@Summary		Add a user's vehicle
//	@Description	adds a vehicle to a user's account
//	@Tags			vehicles
//	@Param			request body models.NewVehiclePostData true "vehicle params"
//	@Success		200 {string} message
//	@Failure		403 {string} message
//	@Router			/auth/vehicles [post]
func (s *Server) AddVehicleToAuthenticatedUserHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)

	vehiclePostData := models.NewVehiclePostData{}

	err := json.NewDecoder(c.Request().Body).Decode(&vehiclePostData)
	if err != nil {
		fmt.Println(err)
		return err
	}

	exists := s.db.CheckUserVehicleExists(claims.ID, vehiclePostData.Registration)
	if exists {
		return c.JSON(http.StatusForbidden, map[string]string{
			"message": "Vehicle already added",
		})
	}

	vehicle, err := utils.FetchVehicleDetails(vehiclePostData)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{
			"message": err.Error(),
		})
	}

	res := s.db.AddNewVehicleToDB(claims.ID, vehicle)

	return c.JSON(http.StatusOK, res)
}

// Get Vehicles godoc
//
//	@Summary		Gets the vehicles for the authenticated user
//	@Description	returns a list of vehicles for the authenticated user
//	@Tags			vehicles
//	@Success		200 {array} models.Vehicle
//	@Failure		403 {string} message
//	@Router			/auth/vehicles [get]
func (s *Server) GetAuthenticatedUserVehiclesHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)

	vehicles, err := s.db.GetAuthenticatedUserVehicles(claims.ID)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, vehicles)
}

// Get Vehicle By ID godoc
//
//	@Summary		Gets a vehicle by ID
//	@Description	returns a vehicle object requested by ID
//	@Tags			vehicles
//	@Param			id	path	string	true	"Vehicle ID"
//	@Success		200 {object} models.Vehicle
//	@Failure		403 {string} message
//	@Router			/auth/vehicles/:id [get]
func (s *Server) GetVehicleByIDHandler(c echo.Context) error {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	vehicle, err := s.db.GetVehicleByID(uuid)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, vehicle)
}

// Get Vehicle By Registration godoc
//
//	@Summary		Gets the authenticated user's vehicle by registration
//	@Description	returns a vehicle object requested by registration
//	@Tags			vehicles
//	@Param			reg	path	string	true	"Registration"
//	@Success		200 {object} models.Vehicle
//	@Failure		403 {string} message
//	@Router			/auth/vehicles/registration/:reg [get]
func (s *Server) GetAuthenticatedUsersVehicleByRegistrationHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)

	reg := c.Param("reg")

	vehicle, err := s.db.GetAuthUserVehicleByRegistration(claims.ID, reg)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, vehicle)
}
