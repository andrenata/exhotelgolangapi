package user

type UserFormatter struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Active int    `json:"active"`
	Token  string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Active: user.Active,
		Token:  token,
	}

	return formatter
}

func FormatProfile(user User) UserFormatter {
	formatter := UserFormatter{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Active: user.Active,
	}
	return formatter
}

func FormatUsers(users []User) []UserFormatter {
	userFormatter := []UserFormatter{}

	for _, user := range users {
		formatter := FormatProfile(user)
		userFormatter = append(userFormatter, formatter)
	}

	return userFormatter
}
