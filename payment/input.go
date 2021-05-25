package payment

type RegisterPaymentInput struct {
	BankName    string `json:"bankname" binding:"required"`
	AccountName string `json:"accountname" binding:"required"`
	BankNumber  string `json:"banknumber" binding:"required"`
	IsActive    int    `json:"isactive" binding:"required"`
}
