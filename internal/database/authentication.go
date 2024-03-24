package database

import (
	"fmt"

	"github.com/ZondaF12/logbook-backend/internal/auth"
	"github.com/ZondaF12/logbook-backend/internal/models"
	"github.com/google/uuid"
)

func (s *service) GetAuthenticationByEmail(email string) models.UserAuth {
	var user models.UserAuth

	sqlStatement := `SELECT * FROM authentication WHERE email=$1`
	row := s.db.QueryRow(sqlStatement, email)

	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		fmt.Println("No user found")
	}
	return user
}

func (s *service) GetAuthenticationByID(id uuid.UUID) models.UserAuth {
	var user models.UserAuth

	sqlStatement := `SELECT * FROM authentication WHERE id=$1`
	row := s.db.QueryRow(sqlStatement, id)

	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		fmt.Println("No user found")
	}
	return user
}

func (s *service) AddAuthenticationToDB(params models.UserAuth) map[string]string {
	exists := s.GetAuthenticationByEmail(params.Email)
	if exists.Email != "" {
		fmt.Println(exists.Email)
		return map[string]string{
			"message": "User already exists",
		}
	}

	hash := auth.HashAndSalt([]byte(params.Password))

	sqlStatement := `INSERT INTO authentication (id, email, password, role) VALUES ($1, $2, $3, $4)`
	_, err := s.db.Exec(sqlStatement, uuid.NewString(), params.Email, hash, "user")
	if err != nil {
		fmt.Println("\nRow not inserted!")
	} else {
		fmt.Println("\nRow inserted successfully!")
	}

	return map[string]string{
		"message": "User Created Successfully",
	}
}
