package Models

import (
	"time"

	"gorm.io/gorm"
)

type Slider struct {
	ID        uint   `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name      string `json:"name"`
	Filename  string `json:"filename"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (b *Slider) TableName() string {
	return "sliders"
}
