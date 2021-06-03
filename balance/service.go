package balance

import "math/rand"

type Service interface {
	TopUpBalance(userId int, input InputTopUp) (BalanceHistory, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) TopUpBalance(userId int, input InputTopUp) (BalanceHistory, error) {
	balanceHistory := BalanceHistory{}
	balanceHistory.UserId = userId
	balanceHistory.PaymentId = input.PaymentId
	balanceHistory.BankSender = input.BankSender
	balanceHistory.NameSender = input.NameSender
	balanceHistory.NumberSender = input.NumberSender
	balanceHistory.Amount = input.Amount
	balanceHistory.Status = 0
	balanceHistory.Code = rand.Intn(999)

	NewBalance, err := s.repository.Save(balanceHistory)
	if err != nil {
		return NewBalance, err
	}

	return NewBalance, nil

}
