package balance

import (
	"cager/user"
)

type InputTopUp struct {
	BankSender   string `json:"bank_sender" binding:"required"`
	NameSender   string `json:"name_sender" binding:"required"`
	NumberSender string `json:"number_sender" binding:"required"`
	Amount       int    `json:"amount" binding:"required"`
	PaymentId    int    `json:"payment_id" binding:"required"`
	User         user.User
}
