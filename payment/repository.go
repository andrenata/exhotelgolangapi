package payment

import "gorm.io/gorm"

type Repository interface {
	Save(payment Payment) (Payment, error)
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
