package balance

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Save(balanceHistory BalanceHistory) (BalanceHistory, error)
	FindByID(ID int) (BalanceHistory, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(balanceHistory BalanceHistory) (BalanceHistory, error) {
	err := r.db.Save(&balanceHistory).Error
	if err != nil {
		return balanceHistory, err
	}
	return balanceHistory, nil
}

func (r *repository) FindByID(ID int) (BalanceHistory, error) {
	var balanceHistory BalanceHistory

	err := r.db.Where("id = ?", ID).Find(&balanceHistory).Error
	if err != nil {
		return balanceHistory, err
	}

	return balanceHistory, nil
}
