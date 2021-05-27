package wallet

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(userId int, wallet Wallet) (Wallet, error)
	// CheckPin(id int, pin string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(userId int, wallet Wallet) (Wallet, error) {
	err := r.db.Create(&wallet).Error
	if err != nil {
		return wallet, err
	}
	return wallet, nil
}
