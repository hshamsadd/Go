package models

import "time"

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email" validate:"required,email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}