package payment

type Service interface {
	RegisterPayment(input RegisterPaymentInput) (Payment, error)
	GetAllPayment() ([]Payment, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterPayment(input RegisterPaymentInput) (Payment, error) {
	payment := Payment{}
	payment.BankName = input.BankName
	payment.AccountName = input.AccountName
	payment.BankNumber = input.BankNumber
	payment.IsActive = input.IsActive

	newPayment, err := s.repository.Save(payment)
	if err != nil {
		return newPayment, err
	}

	return newPayment, nil

}

func (s *service) GetAllPayment() ([]Payment, error) {
	payments, err := s.repository.FindAll()
	if err != nil {
		return payments, err
	}
	return payments, nil
}
