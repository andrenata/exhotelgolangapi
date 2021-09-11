package user

import (
	"time"
)

type User struct {
	ID           int
	Name         string
	Email        string
	Password     string
	PasswordTemp string
	Code         string
	Active       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
