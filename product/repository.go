package product

import "gorm.io/gorm"

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

	err := r.db.Preload("Sliders", "slider.is_primary = 1").Preload("Discounts", "discounts.active = 1").Where("active = ?", active).Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product

	err := r.db.Preload("Sliders", "sliders.is_primary = 1").Preload("Discounts", "discounts.active = 1").Find(&products).Error

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

// ======== SLIDER
func (r *repository) FindAllSlider() ([]Slider, error) {
	var sliders []Slider

	err := r.db.Find(&sliders).Error

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

// ========== DISCOUNT
func (r *repository) FindAllDiscount() ([]Discount, error) {
	var discounts []Discount

	err := r.db.Find(&discounts).Error
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

	err := r.db.Where("product_id = ?", id).Find(&discount).Error

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
