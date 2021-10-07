package settings

import "gorm.io/gorm"

type Repository interface {
	Update(setting Setting) (Setting, error)
	FindById(id int) (Setting, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindById(id int) (Setting, error) {
	var setting Setting
	err := r.db.Where("id = ?", id).Find(&setting).Error
	if err != nil {
		return setting, err
	}
	return setting, nil
}

func (r *repository) Update(setting Setting) (Setting, error) {
	err := r.db.Save(&setting).Error
	if err != nil {
		return setting, err
	}
	return setting, nil
}
