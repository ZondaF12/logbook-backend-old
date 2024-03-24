package models

import "time"

type (
	Vehicles struct {
		Avatar    string    `json:"avatar"`
		Bio       string    `json:"bio"`
		CreatedAt time.Time `json:"created_at"`
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		Public    bool      `json:"public"`
		UserID    string    `json:"user_id"`
		Username  string    `json:"username"`
	}
)
