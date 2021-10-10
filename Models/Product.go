package Models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID               uint               `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name             string             `json:"name"`
	Slug             string             `json:"slug"`
	Thumbnail        string             `json:"thumbnail"`
	Bahan            string             `json:"bahan"`
	Dimensi          string             `json:"dimensi"`
	Price            int                `json:"price"`
	Stock            int                `json:"stock"`
	Active           int                `json:"active"`
	Views            int                `json:"views"`
	Description      string             `json:"description"`
	CategoryRelation []CategoryRelation `gorm:"foreignKey:ProductID"`
	SliderRelation   []SliderRelation   `gorm:"foreignKey:ProductID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}

func (b *Product) TableName() string {
	return "products"
}
