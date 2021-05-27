package wallet

import "time"

type Wallet struct {
	ID              int
	UserId          int
	WalletBalance   int
	WalletTemporary int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
