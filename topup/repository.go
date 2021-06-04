package topup

import (
	"cager/balance"
	"log"
	"os/user"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	WithTrx(*gorm.DB) repository
	IncrementMoney(uint, float64) error
	DecrementMoney(uint, float64) error
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// TRX
// WithTrx enables repository with transaction
func (r repository) WithTrx(trxHandle *gorm.DB) repository {
	if trxHandle == nil {
		log.Print("Database not found")
		return r
	}
	r.db = trxHandle
	return r
}

func (r repository) IncrementMoney(receiver uint, amount float64) error {
	log.Print("[BalanceRepository]...Increment Money")
	return r.db.Model(&user.User{}).Where("id=?", receiver).Update("balance", gorm.Expr("balance + ?", amount)).Error
}

func (r repository) DecrementMoney(giver uint, amount float64) error {
	log.Print("[BalanceRepository]...Decrement Money")
	// return errors.New("something")
	return r.db.Model(&balance.BalanceHistory{}).Where("id=?", giver).Update("status", 1).Error
}
