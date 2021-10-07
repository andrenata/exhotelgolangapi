package product

import (
	"cager/category"
	"cager/helper"

	"fmt"
)

type Service interface {
	FindAllSliderService() ([]Slider, error)
	// UpdateSliderService(id int, input UpdateSliderInput) (Slider, error)
	// UpdateSliderByPostService(idslider int, idproduct int) (Slider, error)
	// FindSliderByIdService(id int) (Slider, error)
	CreateSliderService(name string) (Slider, error)
	DelSliderService(id int) (bool, error)

	// PRODUCT
	FindAllProductService() ([]Product, error)
	FindAllProductBestService() ([]Product, error)
	FindProductById(id int) (Product, error)
	FindProductBySlug(slug string) (bool, error)
	CreateProductService(input CreateProductInput) (Product, error)
	UpdateProductService(input UpdateProductInput) (Product, error)
	DelProductService(id int) (bool, error)
	UpdateProductByActiveService(id int, input UpdateProductByActiveInput) (Product, error)
	CreateProductByName(input CreateProductByName) (Product, error)

	// DISCOUNT
	// FindAllDiscountService() ([]Discount, error)
	CreateDiscountService(input CreateDiscountInput) (Discount, error)
	DelDiscountService(id int) (bool, error)
	// UpdateDiscountService(id int, input CreateDiscountInput) (Discount, error)
	// UpdateDiscountByActiveService(id int, input UpdateDiscountByActiveInput) (Discount, error)

	// SLIDER RELATION
	CheckSliderRelation(product_id int, slider_id int) (bool, error)
	CreateSliderRelationService(input CreateSliderRelationInput) (SliderRelation, error)
	DelSliderRelation(slider_id int, product_id int) (bool, error)
	GetSliderRelationByIDProductService(id int) ([]Slider, error)
	GetSliderRelationByIDService(id int) (SliderRelation, error)

	// CATEGORY RELATION
	CheckCategoryRelation(product_id int, category_id int) (bool, error)
	CreateCategoryRelation(input CreateCategoryRelationInput) (CategoryRelation, error)
	DelCategoryRelation(product_id int, category_id int) (bool, error)
	FindCategoryRelation(id int) ([]category.Category, error)

	ProductPaginationService(url string, pagination *helper.Pagination) helper.ResponsePagination
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAllSliderService() ([]Slider, error) {
	sliders, err := s.repository.FindAllSlider()
	if err != nil {
		return sliders, err
	}

	return sliders, nil
}

func (s *service) CreateSliderService(name string) (Slider, error) {
	slider := Slider{}
	slider.Filename = name

	create, err := s.repository.CreateSlider(slider)
	if err != nil {
		return create, err
	}

	return create, nil

}

func (s *service) UpdateSliderService(id int, input UpdateSliderInput) (Slider, error) {
	slider, err := s.repository.FindSliderById(id)
	if err != nil {
		return slider, err
	}
	slider.Name = input.Name
	slider.Filename = input.Filename

	update, err := s.repository.UpdateSlider(slider)
	if err != nil {
		return update, err
	}
	return update, err
}

func (s *service) FindSliderByIdService(id int) (Slider, error) {
	slider, err := s.repository.FindSliderById(id)
	if err != nil {
		return slider, err
	}

	return slider, nil
}

func (s *service) DelSliderService(id int) (bool, error) {
	slider, err := s.repository.FindSliderById(id)
	if err != nil {
		return false, err
	}

	_, err = s.repository.DelSlider(id, slider)

	if err != nil {
		return false, err
	}

	return true, nil
}

// SLIDER RELATION
func (s *service) CheckSliderRelation(product_id int, slider_id int) (bool, error) {
	check, err := s.repository.CheckSliderRelation(product_id, slider_id)
	if err != nil {
		return false, err
	}

	if check.ID == 0 {
		return true, nil
	}
	return false, nil
}

func (s *service) CreateSliderRelationService(input CreateSliderRelationInput) (SliderRelation, error) {
	sliderRelation := SliderRelation{}
	sliderRelation.ProductID = input.ProductID
	sliderRelation.SliderID = input.SliderID
	create, err := s.repository.CreateSliderRelation(sliderRelation)
	if err != nil {
		return create, err
	}
	return create, nil
}
func (s *service) DelSliderRelation(slider_id int, product_id int) (bool, error) {
	_, err := s.repository.DelSliderRelation(slider_id, product_id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *service) GetSliderRelationByIDProductService(id int) ([]Slider, error) {
	sliderRelation, err := s.repository.GetSliderRelationByProductID(id)
	if err != nil {
		return sliderRelation, err
	}
	return sliderRelation, nil
}

func (s *service) GetSliderRelationByIDService(id int) (SliderRelation, error) {
	sliderRelation, err := s.repository.GetSliderRelationByID(id)
	if err != nil {
		return sliderRelation, err
	}
	return sliderRelation, nil
}

// CATEGORY RELATION
func (s *service) CheckCategoryRelation(product_id int, category_id int) (bool, error) {
	check, err := s.repository.CheckCategoryRelation(product_id, category_id)
	if err != nil {
		return false, err
	}

	if check.ID == 0 {
		return true, nil
	}
	return false, nil
}

func (s *service) CreateCategoryRelation(input CreateCategoryRelationInput) (CategoryRelation, error) {
	categoryRelation := CategoryRelation{}
	categoryRelation.ProductID = input.ProductID
	categoryRelation.CategoryID = input.CategoryID
	relation, err := s.repository.CreateCategoryRelation(categoryRelation)
	if err != nil {
		return relation, err
	}
	return relation, nil
}

func (s *service) DelCategoryRelation(product_id int, category_id int) (bool, error) {
	_, err := s.repository.DelCategoryRelation(product_id, category_id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *service) FindCategoryRelation(id int) ([]category.Category, error) {
	find, err := s.repository.FindCategoryRelation(id)
	if err != nil {
		return find, err
	}
	return find, nil
}

// ============ PRODUCT
func (s *service) CreateProductByName(input CreateProductByName) (Product, error) {
	product := Product{}
	product.Name = input.Name
	product.Slug = input.Slug
	product.Active = 0

	create, err := s.repository.Save(product)
	if err != nil {
		return create, err
	}

	return create, nil
}

func (s *service) FindAllProductService() ([]Product, error) {
	products, err := s.repository.FindAll()
	if err != nil {
		return products, err
	}
	return products, nil
}

func (s *service) FindAllProductBestService() ([]Product, error) {
	Products, err := s.repository.FindAllBest()
	if err != nil {
		return Products, err
	}
	return Products, nil
}

func (s *service) FindProductById(id int) (Product, error) {
	product, err := s.repository.FindById(id)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *service) FindProductBySlug(slug string) (bool, error) {
	product, err := s.repository.FindBySlug(slug)
	if err != nil {
		return false, err
	}

	if product.ID == 0 {
		return true, err
	}

	return false, nil
}

func (s *service) CreateProductService(input CreateProductInput) (Product, error) {
	product := Product{}
	product.Name = input.Name
	product.Slug = input.Slug
	product.Bahan = input.Bahan
	product.Dimensi = input.Dimensi
	product.Price = input.Price
	product.Stock = input.Stock
	product.Description = input.Description
	product.Active = input.Active

	create, err := s.repository.Save(product)
	if err != nil {
		return create, err
	}

	return create, nil
}

func (s *service) UpdateProductService(input UpdateProductInput) (Product, error) {
	product, err := s.repository.FindById(input.ID)
	if err != nil {
		return product, err
	}

	product.Name = input.Name
	product.Slug = input.Slug
	product.Bahan = input.Bahan
	product.Dimensi = input.Dimensi
	product.Price = input.Price
	product.Stock = input.Stock
	product.Description = input.Description
	product.Active = input.Active

	update, err := s.repository.Update(product)
	if err != nil {
		return update, err
	}

	return update, nil
}

func (s *service) DelProductService(id int) (bool, error) {
	product, err := s.repository.FindById(id)
	if err != nil {
		return false, err
	}

	_, err = s.repository.DelProduct(id, product)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *service) UpdateProductByActiveService(id int, input UpdateProductByActiveInput) (Product, error) {
	product, err := s.repository.FindById(id)
	if err != nil {
		return product, err
	}

	product.Active = input.Active
	update, err := s.repository.Update(product)
	if err != nil {
		return update, err
	}
	return update, nil
}

// ========== DISCOUNT
func (s *service) FindAllDiscountService() ([]Discount, error) {
	discounts, err := s.repository.FindAllDiscount()
	if err != nil {
		return discounts, err
	}
	return discounts, nil
}

func (s *service) CreateDiscountService(input CreateDiscountInput) (Discount, error) {
	discount := Discount{}
	discount.Name = input.Name
	discount.Slug = input.Slug
	discount.Persentase = input.Persentase
	discount.Price = input.Price
	discount.Active = input.Active
	discount.ProductID = input.ProductID
	// discount.StartDate = input.StartDate
	// discount.EndDate = input.EndDate

	create, err := s.repository.CreateDiscount(discount)
	if err != nil {
		return create, err
	}

	return create, nil
}

func (s *service) UpdateDiscountService(id int, input CreateDiscountInput) (Discount, error) {
	discount, err := s.repository.FindDiscountById(id)
	if err != nil {
		return discount, err
	}

	discount.Name = input.Name
	discount.Slug = input.Slug
	discount.Persentase = input.Persentase
	discount.Price = input.Price
	discount.Active = input.Active
	discount.ProductID = input.ProductID
	// discount.StartDate = input.StartDate
	// discount.EndDate = input.EndDate

	update, err := s.repository.UpdateDiscount(discount)
	if err != nil {
		return update, err
	}

	return update, nil
}

func (s *service) UpdateDiscountByActiveService(id int, input UpdateDiscountByActiveInput) (Discount, error) {
	discount, err := s.repository.FindDiscountById(id)
	if err != nil {
		return discount, err
	}

	discount.Active = input.Active

	update, err := s.repository.UpdateDiscount(discount)
	if err != nil {
		return update, err
	}

	return update, nil
}

func (s *service) DelDiscountService(id int) (bool, error) {
	discount, err := s.repository.DelDiscount(id)
	if err != nil {
		return false, err
	}

	if discount != true {
		return false, err
	}

	return true, err
}

// PAGINATION
func (s *service) ProductPaginationService(url string, pagination *helper.Pagination) helper.ResponsePagination {

	operationResult, totalPages := s.repository.ProductPagination(pagination)

	if operationResult.Error != nil {
		return helper.ResponsePagination{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*helper.Pagination)

	// get current url path
	urlPath := url

	// search query params
	searchQueryParams := ""

	for _, search := range pagination.Searchs {
		searchQueryParams += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
	}

	// set first & last page pagination response
	data.FirstPage = fmt.Sprintf("%s?size=%d&page=%d&sort=%s", urlPath, pagination.Size, 0, pagination.Sort) + searchQueryParams
	data.LastPage = fmt.Sprintf("%s?size=%d&page=%d&sort=%s", urlPath, pagination.Size, totalPages, pagination.Sort) + searchQueryParams

	if data.Page > 0 {
		// set previous page pagination response
		data.PreviousPage = fmt.Sprintf("%s?size=%d&page=%d&sort=%s", urlPath, pagination.Size, data.Page-1, pagination.Sort) + searchQueryParams
	}

	if data.Page < totalPages {
		// set next page pagination response
		data.NextPage = fmt.Sprintf("%s?size=%d&page=%d&sort=%s", urlPath, pagination.Size, data.Page+1, pagination.Sort) + searchQueryParams
	}

	if data.Page > totalPages {
		// reset previous page
		data.PreviousPage = ""
	}

	return helper.ResponsePagination{Success: true, Data: data}
}
