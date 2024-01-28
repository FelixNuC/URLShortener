package models

import (
	"time"
)

type User struct {
	ID        string    `json:"id"` // UUID para el identificador único del usuario
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Secret    string    `json:"secret"` // Secreto compartido para generar TOTP
}