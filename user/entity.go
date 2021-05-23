package user

import "time"

type User struct {
	ID              int
	Name            string
	Email           string
	Password        string
	Pin             string
	PhoneNumber     string
	TypeVerified    int
	KtpPassport     string
	Address         string
	City            string
	State           string
	Country         string
	Code            string
	PicturePath     string
	IsActive        int
	IsVerified      int
	EmailVerifiedAt time.Time
	PhoneVerifiedAt time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
