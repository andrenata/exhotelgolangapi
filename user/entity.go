package user

import "time"

type User struct {
	ID           int
	Name         string
	Email        string
	Password     string
	TypeVerified string
	KtpPassport  string
	PicturePath  string
	PhoneNumber  string
	Address      string
	City         string
	Province     string
	Country      string
	IsActive     string
	IsVerified   string
	Pin          string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
