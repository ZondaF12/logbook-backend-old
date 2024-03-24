package database

import (
	"fmt"

	"github.com/ZondaF12/logbook-backend/internal/models"
	"github.com/google/uuid"
)

func (s *service) AddNewUserToDB(user models.User) map[string]string {
	sqlStatement := `INSERT INTO users (id, username, name) VALUES ($1, $2, $3)`
	_, err := s.db.Exec(sqlStatement, &user.ID, &user.Username, &user.Name)
	fmt.Println(user)
	if err != nil {
		fmt.Println(err)
		return map[string]string{
			"message": "User already exists",
		}
	}

	return map[string]string{
		"message": "User Created Successfully",
	}
}

func (s *service) GetUsers() []models.User {
	sqlStatement := `SELECT * FROM users`
	rows, err := s.db.Query(sqlStatement)
	if err != nil {
		fmt.Println("No users found")
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Bio, &user.Name, &user.Avatar, &user.Public, &user.CreatedAt)
		if err != nil {
			fmt.Println("Error scanning rows")
		}

		users = append(users, user)
	}
	return users
}

func (s *service) GetUserByID(id uuid.UUID) (models.User, error) {
	var user models.User

	sqlStatement := `SELECT * FROM users WHERE id=$1`
	row := s.db.QueryRow(sqlStatement, id)

	err := row.Scan(&user.ID, &user.Username, &user.Bio, &user.Name, &user.Avatar, &user.Public, &user.CreatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) UpdateUserByID(id uuid.UUID, user models.User) (models.User, error) {
	sqlStatement := `UPDATE users SET bio=$1, name=$2, avatar=$3, public=$4 WHERE id=$5`
	_, err := s.db.Exec(sqlStatement, user.Bio, user.Name, user.Avatar, user.Public, id)
	if err != nil {
		return user, err
	}

	self, err := s.GetUserByID(id)
	if err != nil {
		return models.User{}, err
	}

	return self, nil
}
