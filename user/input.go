package user

type RegisterUserInput struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phonenumber" binding:"required"`
	Address     string `json:"address" binding:"required"`
	City        string `json:"city" binding:"required"`
	State       string `json:"state" binding:"required"`
	Country     string `json:"country" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type TransferInput struct {
	IdGiver    int    `json:"id_giver" binding:"required"`
	IdReciever int    `json:"id_reciever" binding:"required"`
	Amount     int    `json:"amount" binding:"required"`
	Pin        string `json:"pin" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type CheckPhoneInput struct {
	PhoneNumber string `json:"phonenumber" binding:"required"`
}

type ChangeEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type ChangeNameInput struct {
	Name string `json:"name" binding:"required"`
}

type ChangePin struct {
	Pin string `json:"pin" binding:"required"`
}

type ChangePinTemporary struct {
	PinTemporary string `json:"pin_temporary" binding:"required"`
}

type CheckPin struct {
	Pin string `json:"pin" binding:"required"`
}

type CheckPinTemporary struct {
	PinTemporary string `json:"pin_temporary" binding:"required"`
}

type InputChangeNumber struct {
	PhoneNumber string `json:"phonenumber" binding:"required"`
}

// BALANCE
type ChangeBalanceTemp struct {
	BalanceTemporary int `json:"balance_temporary" binding:"required"`
}

type GetBalance struct {
	Token string `json:"token" binding:"required"`
}
