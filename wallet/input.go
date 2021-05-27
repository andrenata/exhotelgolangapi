package wallet

type RegisterWalletInput struct {
	UserId        int `json:"userid" binding:"required"`
	WalletBalance int `json:"wallet" binding:"required"`
	WalletHistory int `json:"wallethistory" binding:"required"`
}
