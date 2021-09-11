package product

import (
	"time"
)

type Product struct {
	ID                int
	Name              string
	Slug              string
	Bahan             string
	Price             int
	Stock             int
	Active            int
	Description       string
	DiscountID        int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Sliders           []Slider
	CategoryRelations []CategoryRelation
}

type Slider struct {
	ID        int
	Name      string
	Filename  string
	ProductID int
	IsPrimary int
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
	CategoryID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
