package database

import (
	"fmt"

	"github.com/ZondaF12/logbook-backend/internal/auth"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
}

func (s *service) GetUserByEmail(email string) User {
	var user User
	sqlStatement := `SELECT * FROM users WHERE email=$1`
	row := s.db.QueryRow(sqlStatement, email)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Name)
	if err != nil {
		fmt.Println("No user found")
	}
	return user
}

func (s *service) AddUserToDB(params User) map[string]string {
	exists := s.GetUserByEmail(params.Email)
	if exists.Email != "" {
		return map[string]string{
			"message": "User already exists",
		}
	}

	hash := auth.HashAndSalt([]byte(params.Password))

	sqlStatement := `INSERT INTO users (id, email, password, name) VALUES ($1, $2, $3, $4)`
	_, err := s.db.Exec(sqlStatement, uuid.NewString(), params.Email, hash, params.Name)
	if err != nil {
		fmt.Println("\nRow not inserted!")
	} else {
		fmt.Println("\nRow inserted successfully!")
	}

	return map[string]string{
		"message": "User Created Successfully",
	}
}