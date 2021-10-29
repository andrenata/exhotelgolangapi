package product

import "time"

type CreateSliderInput struct {
	Name      string `json:"name" binding:"required"`
	Filename  string `json:"filename" binding:"required"`
	ProductID int    `json:"product_id" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateSliderInput struct {
	Name      string `json:"slider" binding:"required"`
	Filename  string `json:"filename" binding:"required"`
	UpdatedAt time.Time
}

type UpdateSliderByPostInput struct {
	ProductId int `json:"product_id" binding:"required"`
	UpdatedAt time.Time
}

type DelSliderInput struct {
	ID int `json:"id" binding:"required"`
}

type FindProductByIdInput struct {
	ID int `json:"id" binding:"required"`
}

// ============= PRODUCT
type PaginationInput struct {
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	Sort      string `json:"sort"`
	Direction string `json:"direction"`
	Active 	int `json:"active"`
	Stock int `json:"stock"`
}

type SearchInput struct {
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	Sort      string `json:"sort"`
	Direction string `json:"direction"`
	Active 	int `json:"active"`
	Stock int `json:"stock"`
	Id int `json:"id"`
	Search string `json:"search"`
}

type PaginationProductCategInput struct {
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	Sort      string `json:"sort"`
	Direction string `json:"direction"`
	Slug string `json:"slug"`
	Stock int `json:"stock"`
}

type UpdateThumbProductInput struct {
	ID int `json:"id"`
	Thumbnail string `json:"thumbnail"`
}

type CreateProductInput struct {
	Name        string `json:"name" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Bahan       string `json:"bahan"`
	Dimensi     string `json:"dimensi"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Active      int    `json:"active"`
	Description string `json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UpdateProductInput struct {
	ID          int    `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Bahan       string `json:"bahan"`
	Dimensi     string `json:"dimensi"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Active      int    `json:"active"`
	Description string `json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UpdateProductByActiveInput struct {
	Active    int `json:"active" binding:"required"`
	UpdatedAt time.Time
}

type CreateProductByName struct {
	Name      string `json:"name" binding:"required"`
	Slug      string `json:"slug" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FindBySlugInput struct {
	Slug string `json:"slug" binding:"required"`
}

type UpdateViewsInput struct {
	ID int `json:"id" binding:"required"`
	Views int `json:"views" binding:"required"`
}

// DISCOUNT
type CreateDiscountInput struct {
	Name       string `json:"name" binding:"required"`
	Slug       string `json:"slug" binding:"required"`
	Persentase int    `json:"persentase"`
	Price      int    `json:"price"`
	ProductID  int    `json:"product_id"`
	// StartDate  time.Time
	// EndDate    time.Time
	Active int `json:"active" binding:"required"`
}

type UpdateDiscountByActiveInput struct {
	Active int `json:"active" binding:"required"`
}

type DelProductInput struct {
	ID int `json:"id" binding:"required"`
}

// SLIDER RELATION
type CreateSliderRelationInput struct {
	ProductID int `json:"product_id" binding:"required"`
	SliderID  int `json:"slider_id" binding:"required"`
}

type IDSliderRelationInput struct {
	ID int `json:"id" binding:"required"`
}

type DelSliderProductInput struct {
	SliderID  int `json:"slider_id" binding:"required"`
	ProductID int `json:"product_id" binding:"required"`
}

// CATEGORY RELATION
type CreateCategoryRelationInput struct {
	CategoryID int `json:"category_id" binding:"required"`
	ProductID  int `json:"product_id" binding:"required"`
}
