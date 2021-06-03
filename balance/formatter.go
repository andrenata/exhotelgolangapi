package balance

type BalanceFormatter struct {
	ID           int                     `json:"id"`
	UserId       int                     `json:"user_id"`
	PaymentId    int                     `json:"payment_id"`
	BankSender   string                  `json:"bank_sender"`
	NameSender   string                  `json:"name_sender"`
	NumberSender string                  `json:"number_sender"`
	Code         int                     `json:"code"`
	Amount       int                     `json:"amount"`
	Status       int                     `json:"status"`
	User         BalanceUserFormatter    `json:"user"`
	Payment      BalancePaymentFormatter `json:"payment"`
}

type BalanceUserFormatter struct {
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

type BalancePaymentFormatter struct {
	BankName    string `json:"bank_name"`
	AccountName string `json:"account_name"`
	BankNumber  string `json:"bank_number"`
}

func FormatBalance(balanceHistory BalanceHistory) BalanceFormatter {
	balanceFormatter := BalanceFormatter{}
	balanceFormatter.ID = balanceHistory.ID
	balanceFormatter.UserId = balanceHistory.UserId
	balanceFormatter.PaymentId = balanceHistory.PaymentId
	balanceFormatter.BankSender = balanceHistory.BankSender
	balanceFormatter.NameSender = balanceHistory.NameSender
	balanceFormatter.NumberSender = balanceHistory.NumberSender
	balanceFormatter.Code = balanceHistory.Code
	balanceFormatter.Amount = balanceHistory.Amount
	balanceFormatter.Status = balanceHistory.Status

	// USER
	user := balanceHistory.User
	balanceUserFormatter := BalanceUserFormatter{}
	balanceUserFormatter.Name = user.Name
	balanceUserFormatter.Balance = user.Balance

	balanceFormatter.User = balanceUserFormatter

	// PAYMENT
	payment := balanceHistory.Payment
	balancePaymentFormatter := BalancePaymentFormatter{}
	balancePaymentFormatter.BankName = payment.BankName
	balancePaymentFormatter.AccountName = payment.AccountName
	balancePaymentFormatter.BankNumber = payment.BankNumber

	balanceFormatter.Payment = balancePaymentFormatter

	return balanceFormatter
}
