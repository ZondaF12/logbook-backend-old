package database

import (
	"fmt"
	"time"

	"github.com/ZondaF12/logbook-backend/internal/models"
	"github.com/google/uuid"
)

func (s *service) CreateSession(user models.User) models.Session {
	session := models.Session{
		UserID:      user.ID,
		Token:       uuid.New(),
		ExpiryDate:  time.Now().Add(time.Hour * 24 * 3),
		CreatedAt:   time.Now(),
		RefreshDate: time.Now().Add(time.Hour * 24),
	}

	sqlStatement := `INSERT INTO sessions (user_id, access_token, expiry_date, created_at, refresh_date) VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.Exec(sqlStatement, session.UserID, session.Token, session.ExpiryDate, session.CreatedAt, session.RefreshDate)
	if err != nil {
		fmt.Println(err)
	}
	return session
}
