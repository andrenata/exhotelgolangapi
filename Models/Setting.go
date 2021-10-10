package Models

import (
	"time"
)

type Setting struct {
	ID          uint   `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Facebook    string `json:"facebook"`
	Instagram   string `json:"instagram"`
	Maps        string `json:"maps"`
	MapsLink  string `json:"maps_link"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Welcome     string `json:"welcome"`
	Banner     string `json:"banner"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (b *Setting) TableName() string {
	return "settings"
}
