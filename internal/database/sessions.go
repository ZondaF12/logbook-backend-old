package database

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	UserID      uuid.UUID `json:"user_id"`
	Token       uuid.UUID `json:"access_token"`
	ExpiryDate  time.Time `json:"expiry_date"`
	CreatedAt   time.Time `json:"created_at"`
	RefreshDate time.Time `json:"refresh_date"`
}

func (s *service) CreateSession(user User) Session {
	session := Session{
		UserID:      user.ID,
		Token:       uuid.New(),
		ExpiryDate:  time.Now().Add(time.Hour * 72),
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
