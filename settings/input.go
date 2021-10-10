package settings

type CreateSettingInput struct {
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Facebook    string `json:"facebook"`
	Instagram   string `json:"instagram"`
	Maps        string `json:"maps"`
	MapsLink    string `json:"maps_link"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Welcome       string `json:"welcome"`
}

type UpdateSettingInput struct {
	ID          int    `json:"id" binding:"required"`
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Facebook    string `json:"facebook"`
	Instagram   string `json:"instagram"`
	Maps        string `json:"maps"`
	MapsLink    string `json:"maps_link"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Welcome       string `json:"welcome"`
}

type IDSettingInput struct {
	ID int `json:"id" binding:"required"`
}
