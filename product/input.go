package product

import "time"

type CreateSliderInput struct {
	Name      string `json:"name" binding:"required"`
	Filename  string `json:"filename" binding:"required"`
	ProductID int    `json:"product_id" binding:"required"`
	IsPrimary int    `json:"is_primary"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateSliderInput struct {
	Name      string `json:"slider" binding:"required"`
	Filename  string `json:"filename" binding:"required"`
	IsPrimary int    `json:"is_primary" binding:"required"`
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
type CreateProductInput struct {
	Name        string `json:"name" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Bahan       string `json:"bahan" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	Active      int    `json:"active" binding:"required"`
	Description string `json:"description" binding:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UpdateProductInput struct {
	ID          int    `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Bahan       string `json:"bahan" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	Active      int    `json:"active" binding:"required"`
	Description string `json:"description" binding:"required"`
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
	Slug string `json:"slug" binding:"requred"`
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
