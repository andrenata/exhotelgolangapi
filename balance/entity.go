package balance

import (
	"cager/payment"
	"cager/user"
	"time"

	"github.com/leekchan/accounting"
)

type BalanceHistory struct {
	ID           int
	UserId       int
	PaymentId    int
	BankSender   string
	NameSender   string
	NumberSender string
	Code         int
	Amount       int
	Status       int
	DeletedAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	User         user.User
	Payment      payment.Payment
}

func (t BalanceHistory) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.Amount)
}
