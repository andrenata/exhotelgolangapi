package payment

import "gorm.io/gorm"

type Repository interface {
	Save(payment Payment) (Payment, error)
	FindAll() ([]Payment, error)
	FindById(id int) (Payment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(payment Payment) (Payment, error) {
	err := r.db.Create(&payment).Error
	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (r *repository) FindAll() ([]Payment, error) {
	var payments []Payment
	err := r.db.Find(&payments).Error
	if err != nil {
		return payments, err
	}
	return payments, nil
}

func (r *repository) FindById(id int) (Payment, error) {
	var payment Payment

	err := r.db.Where("id = ?", id).Find(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil

}
