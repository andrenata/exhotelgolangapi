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
	Views             int
	Description       string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Discounts         []Discount
	SliderRealtions   []SliderRelation
	CategoryRelations []CategoryRelation
}

type Slider struct {
	ID              int
	Name            string
	Filename        string
	ProductID       int
	IsPrimary       int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	SliderRelations []SliderRelation
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
