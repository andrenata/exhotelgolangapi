package product

import (
	"time"
)

type Product struct {
	ID               int                `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name             string             `json:"name"`
	Slug             string             `json:"slug"`
	Bahan            string             `json:"bahan"`
	Thumbnail        string             `json:"thumbnail"`
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
}

type Slider struct {
	ID        int
	Name      string `json:"name"`
	Filename  string `json:"filename"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Discount struct {
	ID         int
	ProductID  int
	Name       string
	Slug       string
	Persentase int
	Price      int
	Active     int
	StartDate  time.Time
	EndDate    time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CategoryRelation struct {
	ID         int
	ProductID  int
	CategoryID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type SliderRelation struct {
	ID        int
	ProductID int
	SliderID  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
