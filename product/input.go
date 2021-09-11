package product

import "time"

type CreateSliderInput struct {
	Name      string `json:"slider" binding:"required"`
	Filename  string `json:"filename" binding:"required"`
	ProductID int    `json:"product_id" binding:"required"`
	IsPrimary int    `json:"is_primary" binding:"required"`
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
	DiscountID  int `json:"discount_id"`
}

type UpdateProductByActiveInput struct {
	Active    int `json:"active" binding:"required"`
	UpdatedAt time.Time
}

// DISCOUNT
type CreateDiscountInput struct {
	Name       string `json:"name" binding:"required"`
	Slug       string `json:"slug" binding:"required"`
	Persentase int    `json:"persentase" binding:"required"`
	Price      int    `json:"price" binding:"required"`
	StartDate  time.Time
	EndDate    time.Time
	Active     int `json:"active" binding:"required"`
}

type UpdateDiscountByActiveInput struct {
	Active int `json:"active" binding:"required"`
}
