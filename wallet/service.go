package wallet

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterWalletUser(input RegisterWalletInput) (Wallet, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterWalletUser(userId int, input RegisterWalletInput) (Wallet, error) {
	wallet, err := s.repository.FindById(userId)
	if err != nil {
		return wallet, err
	}

	Pin, err := bcrypt.GenerateFromPassword([]byte(input.Pin), bcrypt.MinCost)
	if err != nil {
		return wallet, err
	}

	wallet.Pin = string(Pin)
	updatePin, err := s.repository.Update(user)
	if err != nil {
		return updatePin, err
	}

	return updatePin, nil
}
