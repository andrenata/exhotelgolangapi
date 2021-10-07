package settings

type SettingFormatter struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Facebook    string `json:"facebook"`
	Instagram   string `json:"instagram"`
	Maps        string `json:"maps"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
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
	}
	return formatter
}
