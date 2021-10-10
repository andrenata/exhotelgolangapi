package Models

import (
	"time"

	"gorm.io/gorm"
)

type SliderRelation struct {
	ID        uint `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	ProductID int  `json:"product_id"`
	SliderID  int  `json:"slider_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (b *SliderRelation) TableName() string {
	return "slider_relations"
}
