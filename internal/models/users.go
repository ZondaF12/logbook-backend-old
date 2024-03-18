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
		Role     string    `json:"role"`
	}

	SelfUser struct {
		ID    uuid.UUID `json:"id"`
		Email string    `json:"email"`
		Name  string    `json:"name"`
		Role  string    `json:"role"`
	}
)
