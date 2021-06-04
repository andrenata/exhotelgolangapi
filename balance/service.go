package balance

import (
	"cager/payment"
	"cager/user"
	"math/rand"
)

type Service interface {
	TopUpBalance(userId int, input InputTopUp) (BalanceHistory, error)
	FindService(balanceId int) (BalanceHistory, error)
}

type service struct {
	repository     Repository
	userService    user.Service
	paymentService payment.Service
	userRepository user.Repository
}

func NewService(repository Repository, userService user.Service, paymentService payment.Service, userRepository user.Repository) *service {
	return &service{repository, userService, paymentService, userRepository}
}

func (s *service) TopUpBalance(userId int, input InputTopUp) (BalanceHistory, error) {
	user, err := s.userService.GetUserbyId(userId)
	if err != nil {
		return BalanceHistory{}, err
	}

	payment, err := s.paymentService.GetPaymentbyId(input.PaymentId)
	if err != nil {
		return BalanceHistory{}, err
	}

	balanceHistory := BalanceHistory{}
	balanceHistory.UserId = user.ID
	balanceHistory.PaymentId = payment.ID
	balanceHistory.BankSender = input.BankSender
	balanceHistory.NameSender = input.NameSender
	balanceHistory.NumberSender = input.NumberSender
	balanceHistory.Amount = input.Amount
	balanceHistory.Status = 0
	balanceHistory.Code = rand.Intn(999)
	balanceHistory.User = user
	balanceHistory.Payment = payment

	NewBalance, err := s.repository.Save(balanceHistory)
	if err != nil {
		return NewBalance, err
	}

	return NewBalance, nil

}

func (s *service) FindService(balanceId int) (BalanceHistory, error) {
	balanceHistory, err := s.repository.FindByID(balanceId)
	if err != nil {
		return balanceHistory, err
	}
	return balanceHistory, nil
}