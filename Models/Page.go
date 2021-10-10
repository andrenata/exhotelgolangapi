package Models

import (
	"time"

	"gorm.io/gorm"
)

type Page struct {
	ID        uint   `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Description string `json:"description" gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (b *Page) TableName() string {
	return "pages"
}
