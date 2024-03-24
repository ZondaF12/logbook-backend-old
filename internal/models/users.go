package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	UserAuth struct {
		ID       uuid.UUID `json:"id"`
		Email    string    `json:"email"`
		Password string    `json:"password"`
		Role     string    `json:"role"`
	}

	User struct {
		ID        uuid.UUID `json:"id"`
		Username  string    `json:"username"`
		Name      string    `json:"name"`
		Bio       string    `json:"bio"`
		Avatar    string    `json:"avatar"`
		Public    bool      `json:"public"`
		CreatedAt time.Time `json:"created_at"`
	}

	UpdateSelf struct {
		ID        uuid.UUID `json:"id"`
		Username  string    `json:"username"`
		Name      string    `json:"name,omitempty"`
		Bio       string    `json:"bio,omitempty"`
		Avatar    string    `json:"avatar,omitempty"`
		Public    bool      `json:"public,omitempty"`
		CreatedAt time.Time `json:"created_at"`
	}
)
