package pages

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(page Page) (Page, error)
	Update(page Page) (Page, error)
	Delete(id int, page Page) (bool, error)
	FindById(id int) (Page, error)
	FindBySlug(slug string) (Page, error)
	FindAll() ([]Page, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Page, error) {
	var pages []Page
	err := r.db.Order("id desc").Find(&pages).Error
	if err != nil {
		return pages, err
	}
	return pages, nil
}

func (r *repository) FindById(id int) (Page, error) {
	var page Page
	err := r.db.Where("id = ?", id).Find(&page).Error
	if err != nil {
		return page, err
	}
	return page, nil
}

func (r *repository) Save(page Page) (Page, error) {
	err := r.db.Create(&page).Error
	if err != nil {
		return page, err
	}
	return page, nil
}

func (r *repository) Update(page Page) (Page, error) {
	err := r.db.Save(&page).Error
	if err != nil {
		return page, err
	}
	return page, nil
}

func (r *repository) Delete(id int, page Page) (bool, error) {
	err := r.db.Where("id = ?", id).Delete(&page).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *repository) FindBySlug(slug string) (Page, error) {
	var page Page
	err := r.db.Where("slug = ?", slug).Find(&page).Error
	if err != nil {
		return page, err
	}
	return page, nil
}
