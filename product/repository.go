package product

import (
	"cager/category"

	"gorm.io/gorm"
)

type Repository interface {
	Save(product Product) (Product, error)
	Update(product Product) (Product, error)
	FindById(id int) (Product, error)
	// FindByName(name string) (Product, error)
	FindByActive(active int) (Product, error)
	FindAll() ([]Product, error)
	FindProductByCategory(slug string) (Product, error)
	DelProduct(id int, product Product) (bool, error)
	FindBySlug(slug string) (Product, error)

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

func (r *repository) FindAll() ([]Product, error) {
	var products []Product

	err := r.db.Preload("Discounts", "discounts.active = 1").Order("id desc").Find(&products).Error

	if err != nil {
		return products, err
	}
	return products, nil
}

func (r *repository) FindProductByCategory(slug string) (Product, error) {
	var product Product

	err := r.db.Joins("JOIN categories ON categories.product_id = product.id").Where("categories.slug = ?", slug).Find(&product).Error

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

// CATEGORY
