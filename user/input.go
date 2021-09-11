package user

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ChangeEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type ChangeNameInput struct {
	Name string `json:"name" binding:"required"`
}

type ChangePassword struct {
	Password      string `json:"password" binding:"required"`
	PasswordCheck string `json:"password_check" binding:"required"`
}

type DeleteInput struct {
	ID int `json:"id" binding:"required"`
}
