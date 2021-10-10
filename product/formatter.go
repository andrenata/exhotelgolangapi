package product

import "time"

type ProductFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Bahan       string `json:"bahan"`
	Thumbnail       string `json:"thumbnail"`
	Dimensi     string `json:"dimensi"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Active      int    `json:"active"`
	Views       int    `json:"views"`
	Description string `json:"description"`
}

func FormatProduct(product Product) ProductFormatter {
	formatter := ProductFormatter{
		ID:          product.ID,
		Name:        product.Name,
		Slug:        product.Slug,
		Bahan:       product.Bahan,
		Thumbnail:   product.Thumbnail,
		Dimensi:     product.Dimensi,
		Price:       product.Price,
		Stock:       product.Stock,
		Active:      product.Active,
		Views:       product.Views,
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
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Filename string `json:"filename"`
}

func FormatSlider(slider Slider) SliderFormatter {
	formatter := SliderFormatter{
		ID:       slider.ID,
		Name:     slider.Name,
		Filename: slider.Filename,
	}

	return formatter
}

func FormatSliders(sliders []Slider) []SliderFormatter {
	sliderFormatter := []SliderFormatter{}
	for _, slider := range sliders {
		formatter := FormatSlider(slider)
		sliderFormatter = append(sliderFormatter, formatter)
	}
	return sliderFormatter
}

type SliderRelationFormatter struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	SliderID  int `json:"slider_id"`
}

func FormatSliderRelation(sliderRelation SliderRelation) SliderRelationFormatter {
	formatter := SliderRelationFormatter{
		ID:        sliderRelation.ID,
		ProductID: sliderRelation.ProductID,
		SliderID:  sliderRelation.SliderID,
	}

	return formatter
}

func FormatSliderRelations(sliderRelations []SliderRelation) []SliderRelationFormatter {
	SliderRelationFormatter := []SliderRelationFormatter{}
	for _, sliderRelation := range sliderRelations {
		formatter := FormatSliderRelation(sliderRelation)
		SliderRelationFormatter = append(SliderRelationFormatter, formatter)
	}
	return SliderRelationFormatter
}

// CATEGORY RELATION
type CategoryRelationFormatter struct {
	ID         int `json:"id"`
	ProductID  int `json:"product_id"`
	CategoryID int `json:"category_id"`
}

func FormatCategoryRelation(categoryRelation CategoryRelation) CategoryRelationFormatter {
	formatter := CategoryRelationFormatter{
		ID:         categoryRelation.ID,
		ProductID:  categoryRelation.ProductID,
		CategoryID: categoryRelation.CategoryID,
	}

	return formatter
}

func FormatCategoryRelations(categoryRelations []CategoryRelation) []CategoryRelationFormatter {
	CategoryRelationFormatter := []CategoryRelationFormatter{}
	for _, categoryRelation := range categoryRelations {
		formatter := FormatCategoryRelation(categoryRelation)
		CategoryRelationFormatter = append(CategoryRelationFormatter, formatter)
	}
	return CategoryRelationFormatter
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

type RepositoryResult struct {
	Result interface{}
	Error  error
}
