package balance

import "gorm.io/gorm"

type Repository interface {
	Save(balanceHistory BalanceHistory) (BalanceHistory, error)
}

type repository struct {
	db *gorm.DB
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
