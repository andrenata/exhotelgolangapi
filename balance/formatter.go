package balance

type BalanceFormatter struct {
	ID           int    `json:"id"`
	UserId       int    `json:"user_id"`
	PaymentId    int    `json:"payment_id"`
	BankSender   string `json:"bank_sender"`
	NameSender   string `json:"name_sender"`
	NumberSender string `json:"number_sender"`
	Code         int    `json:"code"`
	Amount       int    `json:"amount"`
	Status       int    `json:"status"`
}

func FormatBalance(balanceHistory BalanceHistory) BalanceFormatter {
	formatter := BalanceFormatter{
		ID:           balanceHistory.ID,
		UserId:       balanceHistory.UserId,
		PaymentId:    balanceHistory.PaymentId,
		BankSender:   balanceHistory.BankSender,
		NameSender:   balanceHistory.NameSender,
		NumberSender: balanceHistory.NumberSender,
		Code:         balanceHistory.Code,
		Amount:       balanceHistory.Amount,
		Status:       balanceHistory.Status,
	}
	return formatter
}
