package Models

import (
	"time"

	"gorm.io/gorm"
)

type CategoryRelation struct {
	ID         uint `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	ProductID  int  `json:"product_id"`
	CategoryID int  `json:"category_id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

func (b *CategoryRelation) TableName() string {
	return "category_relations"
}
