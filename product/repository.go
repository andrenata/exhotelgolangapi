package product

import (
	"cager/Models"
	"cager/category"
	"cager/helper"

	"fmt"
	"math"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	Save(product Product) (Product, error)
	Update(product Product) (Product, error)
	FindById(id int) (Product, error)
	// FindByName(name string) (Product, error)
	FindByActive(active int) (Product, error)
	FindAll(input PaginationInput) ([]Product, int64)
	SearchAll(input SearchInput) ([]Product, int64)
	FindAllBest() ([]Product, error)
	FindProductByCategory(slug string) (Product, error)
	DelProduct(id int, product Product) (bool, error)
	FindBySlug(slug string) (Product, error)
	GetSliderRelationByProductSlug(slug string) ([]Slider, error)

	// SLIDER
	FindAllSlider() ([]Slider, error)
	CreateSlider(slider Slider) (Slider, error)
	UpdateSlider(slider Slider) (Slider, error)
	FindSliderByProduct(id int) (Slider, error)
	FindSliderById(id int) (Slider, error)
	DelSlider(id int, slider Slider) (bool, error)

	// DISCOUNT
	FindAllDiscount() ([]Discount, error)
	CreateDiscount(discount Discount) (Discount, error)
	UpdateDiscount(discount Discount) (Discount, error)
	FindDiscountByProduct(id int) (Discount, error)
	FindDiscountById(id int) (Discount, error)
	DelDiscount(id int) (bool, error)

	// SLIDER RELATION
	CheckSliderRelation(product_id int, slider_id int) (SliderRelation, error)
	CreateSliderRelation(sliderRelation SliderRelation) (SliderRelation, error)
	GetSliderRelationByProductID(id int) ([]Slider, error)
	GetSliderRelationByID(id int) (SliderRelation, error)
	DelSliderRelation(slider_id int, product_id int) (bool, error)

	// CATEGORY RELATION
	CreateCategoryRelation(categoryRelation CategoryRelation) (CategoryRelation, error)
	DelCategoryRelation(product_id int, category_id int) (bool, error)
	FindCategoryRelation(id int) ([]category.Category, error)
	CheckCategoryRelation(product_id int, category_id int) (CategoryRelation, error)

	ProductPagination(pagination *helper.Pagination) (RepositoryResult, int)

	//FRONT END
	FrontFindProductByCateg(input PaginationProductCategInput) ([]Product, int64)

	// Category
	// FindCategoryByProduct(id int) (Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(product Product) (Product, error) {
	err := r.db.Create(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Update(product Product) (Product, error) {
	err := r.db.Save(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FindById(id int) (Product, error) {
	var product Product

	err := r.db.Where("id = ?", id).Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FindBySlug(slug string) (Product, error) {
	var product Product

	err := r.db.Where("slug = ?", slug).Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FindByActive(active int) (Product, error) {
	var product Product

	err := r.db.Preload("Discounts", "discounts.active = 1").Where("active = ?", active).Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FrontFindProductByCateg(input PaginationProductCategInput) ([]Product, int64) {
	var products []Product
	totalRows := 0
	total := int64(totalRows)

	//query
	offset := input.Page * input.Size
	find := r.db.Limit(input.Size).Offset(offset).Order(input.Sort+" "+input.Direction)
	if input.Stock == 1 {
		find = find.Joins("JOIN category_relations ON products.id = category_relations.product_id").
			Joins("JOIN categories ON category_relations.category_id = categories.id").
			Where("categories.slug = ?", input.Slug).
			Where("products.active = ?", 1).
			Where("products.stock > ?", 0).
			Find(&products)
	}
	if input.Stock == 0 {
		find = find.Joins("JOIN category_relations ON products.id = category_relations.product_id").
			Joins("JOIN categories ON category_relations.category_id = categories.id").
			Where("categories.slug = ?", input.Slug).
			Where("products.active = ?", 1).
			Where("products.stock = ?", 0).
			Find(&products)
	}
	err := find.Error
	if err != nil {
		return products, total
	}

	data := products

	// count all data
	if input.Stock == 1 {
		err = r.db.Joins("JOIN category_relations ON products.id = category_relations.product_id").
			Joins("JOIN categories ON category_relations.category_id = categories.id").
			Where("categories.slug = ?", input.Slug).
			Where("products.active = ?", 1).
			Where("products.stock > ?", 0).
			Find(&products).Count(&total).Error
	}

	if input.Stock == 0 {
		err = r.db.Joins("JOIN category_relations ON products.id = category_relations.product_id").
			Joins("JOIN categories ON category_relations.category_id = categories.id").
			Where("categories.slug = ?", input.Slug).
			Where("products.active = ?", 1).
			Where("products.stock = ?", 0).
			Find(&products).Count(&total).Error
	}

	if err != nil {
		return data, total
	}

	return data, total
}

func (r *repository) FindAll(input PaginationInput) ([]Product, int64) {
	var products []Product
	totalRows := 0
	total := int64(totalRows)

	//query
	offset := input.Page * input.Size
	find := r.db.Limit(input.Size).Offset(offset).Order(input.Sort+" "+input.Direction)

	if input.Active == 1 {
		if input.Stock == 1 {
			find = find.Where("active = ?", 1).
				Where("stock > ?", 0).
				Find(&products)
		}

		if input.Stock == 0 {
			find = find.Where("active = ?", 1).
				Where("stock = ?", 0).
				Find(&products)
		}

	}
	if input.Active == 0 {
		find = find.Find(&products)
	}
	err := find.Error
	if err != nil {
		return products, total
	}

	data := products

	// count all data
	if input.Active == 1 {
		if input.Stock == 1 {
			err = r.db.Where("active = ?", 1).
				Where("stock > ?", 0).
				Find(&products).Count(&total).Error
		}

		if input.Stock == 0 {
			err = r.db.Where("active = ?", 1).
				Where("stock = ?", 0).
				Find(&products).Count(&total).Error
		}
	}
	if input.Active == 0 {
		err = r.db.Find(&products).Count(&total).Error
	}
	if err != nil {
		return data, total
	}

	return data, total
}

func (r *repository) SearchAll(input SearchInput) ([]Product, int64) {
	var products []Product
	totalRows := 0
	total := int64(totalRows)

	//query
	offset := input.Page * input.Size
	find := r.db.Limit(input.Size).Offset(offset).Order(input.Sort+" "+input.Direction)

	if input.Id > 0 {
		find = find.Where("active = ?", 1).
			Where("stock > ?", 0).
			Where("id = ?", input.Id).
			Find(&products)
	}

	if input.Search != "" {
		find = find.Where("active = ?", 1).
			Where("stock > ?", 0).
			Where("name LIKE ?", "%"+input.Search+"%").
			Or("slug LIKE ?", "%"+input.Search+"%").
			Or("description LIKE ?", "%"+input.Search+"%").
			Find(&products)
	}


	if input.Active == 0 {
		find = find.Find(&products)
	}
	err := find.Error
	if err != nil {
		return products, total
	}

	data := products

	// count all data
	if input.Id > 0 {
		err = r.db.Where("active = ?", 1).
			Where("stock > ?", 0).
			Where("id = ?", input.Id).
			Find(&products).Count(&total).Error
	}

	if input.Search != "" {
		err = r.db.Where("active = ?", 1).
			Where("stock > ?", 0).
			Where("name LIKE ?", "%"+input.Search+"%").
			Or("slug LIKE ?", "%"+input.Search+"%").
			Or("description LIKE ?", "%"+input.Search+"%").
			Find(&products).Count(&total).Error
	}


	if input.Active == 0 {
		err = r.db.Find(&products).Count(&total).Error
	}
	if err != nil {
		return data, total
	}

	return data, total
}

func (r *repository) FindAllBest() ([]Product, error) {
	var products []Product

	err := r.db.Where("active = 1").Order("views desc").Find(&products).Error

	if err != nil {
		return products, err
	}

	return products, nil
}



func (r *repository) FindProductByCategory(slug string) (Product, error) {
	var product Product

	err := r.db.Joins("JOIN categories ON categories.product_id = product.id").
		Where("categories.slug = ?", slug).Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) DelProduct(id int, product Product) (bool, error) {
	err := r.db.Where("id = ?", id).Delete(product).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// CATEGORY RELATION
func (r *repository) CreateCategoryRelation(categoryRelation CategoryRelation) (CategoryRelation, error) {

	err := r.db.Create(&categoryRelation).Error
	if err != nil {
		return categoryRelation, err
	}
	return categoryRelation, nil
}

func (r *repository) DelCategoryRelation(product_id int, category_id int) (bool, error) {
	var categoryRelation CategoryRelation
	err := r.db.Where("product_id = ?", product_id).
		Where("category_id = ?", category_id).
		Delete(&categoryRelation).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *repository) FindCategoryRelation(id int) ([]category.Category, error) {
	var category []category.Category
	err := r.db.Table("categories").
		Joins("JOIN category_relations ON categories.id = category_relations.category_id").
		Where("category_relations.product_id = ?", id).
		Order("categories.id desc").
		Find(&category).Error

	if err != nil {
		return category, err
	}
	return category, nil
}

func (r *repository) CheckCategoryRelation(product_id int, category_id int) (CategoryRelation, error) {
	var categoryRelation CategoryRelation
	err := r.db.Where("product_id = ?", product_id).
		Where("category_id = ?", category_id).
		Find(&categoryRelation).Error

	if err != nil {
		return categoryRelation, err
	}
	return categoryRelation, nil
}

// ======== SLIDER
func (r *repository) FindAllSlider() ([]Slider, error) {
	var sliders []Slider

	err := r.db.Order("id desc").Find(&sliders).Error

	if err != nil {
		return sliders, err
	}
	return sliders, nil
}

func (r *repository) CreateSlider(slider Slider) (Slider, error) {
	err := r.db.Create(&slider).Error

	if err != nil {
		return slider, err
	}

	return slider, nil
}

func (r *repository) UpdateSlider(slider Slider) (Slider, error) {
	err := r.db.Save(&slider).Error

	if err != nil {
		return slider, err
	}

	return slider, nil
}

func (r *repository) FindSliderByProduct(id int) (Slider, error) {
	var slider Slider

	err := r.db.Where("product_id = ?", id).Find(&slider).Error

	if err != nil {
		return slider, err
	}

	return slider, nil
}

func (r *repository) FindSliderById(id int) (Slider, error) {
	var slider Slider

	err := r.db.Where("id = ?", id).Find(&slider).Error

	if err != nil {
		return slider, err
	}

	return slider, nil
}

func (r *repository) DelSlider(id int, slider Slider) (bool, error) {
	err := r.db.Where("id = ?", id).Delete(&slider).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

// SLIDER RELATION
func (r *repository) CheckSliderRelation(product_id int, slider_id int) (SliderRelation, error) {
	var sliderRelation SliderRelation
	err := r.db.Where("product_id = ?", product_id).
		Where("slider_id = ?", slider_id).
		Order("id desc").
		Find(&sliderRelation).Error

	if err != nil {
		return sliderRelation, err
	}
	return sliderRelation, nil
}

func (r *repository) CreateSliderRelation(sliderRelation SliderRelation) (SliderRelation, error) {
	err := r.db.Create(&sliderRelation).Error
	if err != nil {
		return sliderRelation, err
	}

	return sliderRelation, nil
}

func (r *repository) GetSliderRelationByProductID(id int) ([]Slider, error) {
	var sliders []Slider

	err := r.db.Table("sliders").
		Joins("JOIN slider_relations ON sliders.id = slider_relations.slider_id").
		Where("slider_relations.product_id = ?", id).
		Order("sliders.id desc").
		Find(&sliders).Error

	if err != nil {
		return sliders, err
	}

	return sliders, nil
}

func (r *repository) GetSliderRelationByProductSlug(slug string) ([]Slider, error) {
	var sliders []Slider

	err := r.db.Table("sliders").
		Joins("JOIN slider_relations ON sliders.id = slider_relations.slider_id").
		Joins("JOIN products ON slider_relations.product_id = products.id").
		Where("products.slug = ?", slug).
		Order("sliders.id desc").
		Find(&sliders).Error

	if err != nil {
		return sliders, err
	}

	return sliders, nil
}

func (r *repository) GetSliderRelationByID(id int) (SliderRelation, error) {
	var sliderRelation SliderRelation

	err := r.db.Where("id = ?", id).Find(&sliderRelation).Error

	if err != nil {
		return sliderRelation, err
	}

	return sliderRelation, nil
}

func (r *repository) DelSliderRelation(slider_id int, product_id int) (bool, error) {

	var sliderRelation SliderRelation

	err := r.db.Where("slider_id = ?", slider_id).
		Where("product_id = ?", product_id).
		Delete(&sliderRelation).Error

	// err := r.db.Table("sliders").
	// 	Joins("JOIN slider_relations ON slider_relations.slider_id = sliders.id").
	// 	Where("sliders.id = ?", id).
	// 	Where("")
	// 	Delete(&sliderRelation).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

// ========== DISCOUNT
func (r *repository) FindAllDiscount() ([]Discount, error) {
	var discounts []Discount

	err := r.db.Find(&discounts).Order("id desc").Error
	if err != nil {
		return discounts, err
	}

	return discounts, nil
}

func (r *repository) CreateDiscount(discount Discount) (Discount, error) {
	err := r.db.Create(&discount).Error

	if err != nil {
		return discount, err
	}

	return discount, nil
}

func (r *repository) UpdateDiscount(discount Discount) (Discount, error) {
	err := r.db.Save(&discount).Error

	if err != nil {
		return discount, err
	}

	return discount, nil
}

func (r *repository) FindDiscountByProduct(id int) (Discount, error) {
	var discount Discount

	err := r.db.Where("product_id = ?", id).Find(&discount).Order("id desc").Error

	if err != nil {
		return discount, err
	}

	return discount, nil
}

func (r *repository) FindDiscountById(id int) (Discount, error) {
	var discount Discount

	err := r.db.Where("id = ?", id).Find(&discount).Error

	if err != nil {
		return discount, err
	}

	return discount, nil
}

func (r *repository) DelDiscount(id int) (bool, error) {
	var discount Discount

	err := r.db.Where("id = ?", id).Delete(&discount).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

// PAGINATION

func (r *repository) ProductPagination(pagination *helper.Pagination) (RepositoryResult, int) {
	var products []Models.Product

	var totalRows int64

	totalRows, totalPages, fromRow, toRow := 0, 0, 0, 0

	offset := pagination.Page * pagination.Size

	// get data with limit, offset & order
	find := r.db.Limit(pagination.Size).Offset(offset).Order(pagination.Sort)

	// generate where query
	searchs := pagination.Searchs

	if searchs != nil {
		for _, value := range searchs {
			column := value.Column
			action := value.Action
			query := value.Query

			switch action {
			case "equals":
				whereQuery := fmt.Sprintf("%s = ?", column)
				find = find.Where(whereQuery, query)
				break
			case "contains":
				whereQuery := fmt.Sprintf("%s LIKE ?", column)
				find = find.Where(whereQuery, "%"+query+"%")
				break
			case "in":
				whereQuery := fmt.Sprintf("%s IN (?)", column)
				queryArray := strings.Split(query, ",")
				find = find.Where(whereQuery, queryArray)
				break

			}
		}
	}

	find = find.Find(&products)
	errFind := find.Error
	if find != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}
	pagination.Rows = products

	// count all data
	totalRows = int64(totalRows)
	errCount := r.db.Find(&products).Count(&totalRows).Error
	if errCount != nil {
		return RepositoryResult{Error: errCount}, totalPages
	}

	pagination.TotalRows = totalRows

	// calculate total pages
	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Size))) - 1

	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = pagination.Size
	} else {
		if pagination.Page <= totalPages {
			// calculate from & to row
			fromRow = pagination.Page*pagination.Size + 1
			toRow = (pagination.Page + 1) * pagination.Size
		}
	}

	if toRow > int(totalRows) {
		// set to row with total rows
		toRow = int(totalRows)
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return RepositoryResult{Result: pagination}, totalPages
}
