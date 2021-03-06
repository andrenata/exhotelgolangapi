package category

import "gorm.io/gorm"

type Repository interface {
	Save(category Category) (Category, error)
	Update(category Category) (Category, error)
	FindById(id int) (Category, error)
	FindBySlug(slug string) (Category, error)
	Delete(id int, category Category) (bool, error)
	FindAll() ([]Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(category Category) (Category, error) {
	err := r.db.Create(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) Update(category Category) (Category, error) {
	err := r.db.Save(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) FindAll() ([]Category, error) {
	var categories []Category

	err := r.db.Find(&categories).Error

	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (r *repository) FindById(id int) (Category, error) {
	var category Category

	err := r.db.Where("id = ?", id).Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil

}
func (r *repository) FindBySlug(slug string) (Category, error) {
	var category Category

	err := r.db.Where("slug = ?", slug).Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil

}

func (r *repository) Delete(id int, category Category) (bool, error) {
	err := r.db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
