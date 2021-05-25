package payment

import "time"

type Payment struct {
	ID          int
	BankName    string
	AccountName string
	BankNumber  string
	IsActive    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
