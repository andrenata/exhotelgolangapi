package user

import "time"

type UserFormatter struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	PhoneNumber      string    `json:"phonenumber"`
	TypeVerified     int       `json:"typeverified"`
	KtpPassport      string    `json:"ktppassport"`
	Address          string    `json:"address"`
	City             string    `json:"city"`
	State            string    `json:"state"`
	Country          string    `json:"country"`
	Balance          int       `json:"balance"`
	Token            string    `json:"token"`
	BalanceTemporary int       `json:"balancetemporary"`
	PicturePath      string    `json:"picturepath"`
	IsActive         int       `json:"isactive"`
	IsVerified       int       `json:"isverified"`
	EmailVerifiedAt  time.Time `json:"emailverifiedat"`
	PhoneVerifiedAt  time.Time `json:"phoneverifiedat"`
}

type BalanceFormatter struct {
	Balance int `json:"balance"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Balance:     user.Balance,
		Token:       token,
	}

	return formatter
}

func FormatProfile(user User) UserFormatter {
	formatter := UserFormatter{
		ID:               user.ID,
		Name:             user.Name,
		Email:            user.Email,
		PhoneNumber:      user.PhoneNumber,
		TypeVerified:     user.TypeVerified,
		KtpPassport:      user.KtpPassport,
		Address:          user.Address,
		City:             user.City,
		State:            user.State,
		Country:          user.Country,
		Balance:          user.Balance,
		BalanceTemporary: user.BalanceTemporary,
		PicturePath:      user.PicturePath,
		IsActive:         user.IsActive,
		IsVerified:       user.IsVerified,
		EmailVerifiedAt:  user.EmailVerifiedAt,
		PhoneVerifiedAt:  user.PhoneVerifiedAt,
	}
	return formatter
}

func FormatBalance(user User) BalanceFormatter {
	formatter := BalanceFormatter{
		Balance: user.Balance,
	}
	return formatter
}
