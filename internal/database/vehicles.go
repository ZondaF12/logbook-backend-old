package database

import (
	"github.com/ZondaF12/logbook-backend/internal/models"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (s *service) CheckUserVehicleExists(userId uuid.UUID, registration string) bool {
	sqlStatement := `SELECT EXISTS(SELECT 1 FROM vehicles WHERE user_id=$1 AND registration=$2)`
	row := s.db.QueryRow(sqlStatement, userId, registration)

	var exists bool
	row.Scan(&exists)

	return exists
}

func (s *service) AddNewVehicleToDB(userId uuid.UUID, vehicle models.NewVehicle) map[string]string {
	sqlStatement := `INSERT INTO vehicles (id, user_id, registration, color, description, engine_size, images, insurance_date, make, model, mot_date, nickname, registered, service_date, tax_date, year) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`
	_, err := s.db.Exec(sqlStatement, uuid.NewString(), userId, vehicle.Registration, vehicle.Color, vehicle.Description, vehicle.EngineSize, vehicle.Images, vehicle.InsuranceDate, vehicle.Make, vehicle.Model, vehicle.MotDate, vehicle.Nickname, vehicle.Registered, vehicle.ServiceDate, vehicle.TaxDate, vehicle.Year)

	if err != nil {
		return map[string]string{
			"message": err.Error(),
		}
	}

	return map[string]string{
		"message": "Vehicle added successfully",
	}
}

func (s *service) GetAuthenticatedUserVehicles(userId uuid.UUID) ([]models.Vehicle, error) {
	sqlStatement := `SELECT * FROM vehicles WHERE user_id=$1`
	rows, err := s.db.Query(sqlStatement, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicles []models.Vehicle
	for rows.Next() {
		var vehicle models.Vehicle
		err := rows.Scan(&vehicle.ID, &vehicle.UserID, &vehicle.Registration, &vehicle.Color, &vehicle.Description, &vehicle.EngineSize, pq.Array(&vehicle.Images), &vehicle.InsuranceDate, &vehicle.Make, &vehicle.Model, &vehicle.MotDate, &vehicle.Nickname, &vehicle.Registered, &vehicle.ServiceDate, &vehicle.TaxDate, &vehicle.Year)
		if err != nil {
			return nil, err
		}

		vehicles = append(vehicles, vehicle)
	}
	return vehicles, nil
}

func (s *service) GetVehicleByID(vehicleId uuid.UUID) (models.Vehicle, error) {
	sqlStatement := `SELECT * FROM vehicles WHERE id=$1`
	row := s.db.QueryRow(sqlStatement, vehicleId)

	var vehicle models.Vehicle
	err := row.Scan(&vehicle.ID, &vehicle.UserID, &vehicle.Registration, &vehicle.Color, &vehicle.Description, &vehicle.EngineSize, pq.Array(&vehicle.Images), &vehicle.InsuranceDate, &vehicle.Make, &vehicle.Model, &vehicle.MotDate, &vehicle.Nickname, &vehicle.Registered, &vehicle.ServiceDate, &vehicle.TaxDate, &vehicle.Year)
	if err != nil {
		return vehicle, err
	}

	return vehicle, nil
}

func (s *service) GetAuthUserVehicleByRegistration(userId uuid.UUID, registration string) (models.Vehicle, error) {
	sqlStatement := `SELECT * FROM vehicles WHERE user_id=$1 AND registration=$2`
	row := s.db.QueryRow(sqlStatement, userId, registration)

	var vehicle models.Vehicle
	err := row.Scan(&vehicle.ID, &vehicle.UserID, &vehicle.Registration, &vehicle.Color, &vehicle.Description, &vehicle.EngineSize, pq.Array(&vehicle.Images), &vehicle.InsuranceDate, &vehicle.Make, &vehicle.Model, &vehicle.MotDate, &vehicle.Nickname, &vehicle.Registered, &vehicle.ServiceDate, &vehicle.TaxDate, &vehicle.Year)
	if err != nil {
		return vehicle, err
	}

	return vehicle, nil
}
