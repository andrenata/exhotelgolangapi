package product

import "time"

type ProductFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Bahan       string `json:"bahan"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Active      int    `json:"active"`
	Description string `json:"description"`
}

func FormatProduct(product Product) ProductFormatter {
	formatter := ProductFormatter{
		ID:          product.ID,
		Name:        product.Name,
		Slug:        product.Slug,
		Bahan:       product.Bahan,
		Price:       product.Price,
		Stock:       product.Stock,
		Active:      product.Active,
		Description: product.Description,
	}
	return formatter
}

func FormatProducts(products []Product) []ProductFormatter {

	productFormatter := []ProductFormatter{}

	for _, product := range products {
		formatter := FormatProduct(product)
		productFormatter = append(productFormatter, formatter)
	}

	return productFormatter
}

// SLIDER
type SliderFormatter struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Filename  string `json:"filename"`
	IsPrimary int    `json:"is_primary"`
}

func FormatSlider(slider Slider) SliderFormatter {
	formatter := SliderFormatter{
		ID:        slider.ID,
		Name:      slider.Name,
		Filename:  slider.Filename,
		IsPrimary: slider.IsPrimary,
	}

	return formatter
}

// DISCOUNT
type DiscountFormatter struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Slug       string    `json:"slug"`
	Persentase int       `json:"persentase"`
	Price      int       `json:"price"`
	Active     int       `json:"active"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	ProductID  int       `json:"product_id"`
}

func FormatDiscount(discount Discount) DiscountFormatter {
	formatter := DiscountFormatter{
		Name:       discount.Name,
		Slug:       discount.Slug,
		Persentase: discount.Persentase,
		Price:      discount.Price,
		Active:     discount.Active,
		// StartDate:  discount.StartDate,
		// EndDate:    discount.EndDate,
		ProductID: discount.ProductID,
	}

	return formatter
}
