package pulsa

type service struct {
}

type Service interface {
	// FindProductByBrand(input FindProductInput) (string, error)
}

func NewService() *service {
	return &service{}
}

// func (s *service) FindProductByBrand(input FindProductInput) (string, error) {

// }
