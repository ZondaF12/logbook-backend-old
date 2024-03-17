package models

import (
	"github.com/google/uuid"
)

type (
	User struct {
		ID       uuid.UUID `json:"id"`
		Email    string    `json:"email"`
		Name     string    `json:"name"`
		Password string    `json:"password"`
	}

	SelfUser struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	Self struct {
		SelfUser
		Session
	}
)
