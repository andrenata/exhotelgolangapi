package settings

type CreateSettingInput struct {
	Title       string `json:"title" binding:"required"`
	Keyword     string `json:"keyword" binding:"required"`
	Description string `json:"description" binding:"required"`
	Facebook    string `json:"facebook" binding:"required"`
	Instagram   string `json:"instagram" binding:"required"`
	Maps        string `json:"maps" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
}

type UpdateSettingInput struct {
	ID          int    `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Keyword     string `json:"keyword" binding:"required"`
	Description string `json:"description" binding:"required"`
	Facebook    string `json:"facebook" binding:"required"`
	Instagram   string `json:"instagram" binding:"required"`
	Maps        string `json:"maps" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
}

type IDSettingInput struct {
	ID int `json:"id" binding:"required"`
}
