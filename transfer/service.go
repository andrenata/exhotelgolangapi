package transfer

import (
	"gorm.io/gorm"
)

type Service interface {
	WithTrx(*gorm.DB) service
	IncrementMoney(uint, float64) error
	DecrementMoney(uint, float64) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// WithTrx enables repository with transaction
func (s service) WithTrx(trxHandle *gorm.DB) service {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

func (s service) IncrementMoney(reciever uint, amount float64) error {
	return s.repository.IncrementMoney(reciever, amount)
}
func (s service) DecrementMoney(giver uint, amount float64) error {
	return s.repository.DecrementMoney(giver, amount)
}
