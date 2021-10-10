package settings

type SettingFormatter struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Facebook    string `json:"facebook"`
	Instagram   string `json:"instagram"`
	Maps        string `json:"maps"`
	MapsLink        string `json:"maps_link"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Welcome       string `json:"welcome"`
	Banner       string `json:"banner"`
}

func FormatSetting(setting Setting) SettingFormatter {
	formatter := SettingFormatter{
		ID:          setting.ID,
		Title:       setting.Title,
		Keyword:     setting.Keyword,
		Description: setting.Description,
		Facebook:    setting.Facebook,
		Instagram:   setting.Instagram,
		Maps:        setting.Maps,
		Address:     setting.Address,
		Phone:       setting.Phone,
		MapsLink:       setting.MapsLink,
		Email:       setting.Email,
		Welcome:       setting.Welcome,
		Banner:       setting.Banner,
	}
	return formatter
}
