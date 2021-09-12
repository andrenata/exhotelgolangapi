package product

type Service interface {
	// FindAllSliderService() ([]Slider, error)
	CreateSliderService(input CreateSliderInput) (Slider, error)
	// UpdateSliderService(id int, input UpdateSliderInput) (Slider, error)
	// UpdateSliderByPostService(idslider int, idproduct int) (Slider, error)
	// FindSliderByIdService(id int) (Slider, error)
	DelSliderService(id int) (bool, error)

	// PRODUCT
	FindAllProductService() ([]Product, error)
	FindProductById(id int) (Product, error)
	FindProductBySlug(slug string) (bool, error)
	CreateProductService(input CreateProductInput) (Product, error)
	UpdateProductService(input UpdateProductInput) (Product, error)
	DelProductService(id int) (bool, error)
	UpdateProductByActiveService(id int, input UpdateProductByActiveInput) (Product, error)
	CreateProductByName(input CreateProductByName) (Product, error)

	// DISCOUNT
	// FindAllDiscountService() ([]Discount, error)
	CreateDiscountService(input CreateDiscountInput) (Discount, error)
	DelDiscountService(id int) (bool, error)
	// UpdateDiscountService(id int, input CreateDiscountInput) (Discount, error)
	// UpdateDiscountByActiveService(id int, input UpdateDiscountByActiveInput) (Discount, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAllSliderService() ([]Slider, error) {
	sliders, err := s.repository.FindAllSlider()
	if err != nil {
		return sliders, err
	}

	return sliders, nil
}

func (s *service) CreateSliderService(input CreateSliderInput) (Slider, error) {
	slider := Slider{}
	slider.Name = input.Name
	slider.Filename = input.Filename
	slider.IsPrimary = input.IsPrimary
	slider.ProductID = input.ProductID

	create, err := s.repository.CreateSlider(slider)
	if err != nil {
		return create, err
	}

	return create, nil

}

func (s *service) UpdateSliderService(id int, input UpdateSliderInput) (Slider, error) {
	slider, err := s.repository.FindSliderById(id)
	if err != nil {
		return slider, err
	}
	slider.Name = input.Name
	slider.Filename = input.Filename
	slider.IsPrimary = input.IsPrimary

	update, err := s.repository.UpdateSlider(slider)
	if err != nil {
		return update, err
	}
	return update, err
}

func (s *service) FindSliderByIdService(id int) (Slider, error) {
	slider, err := s.repository.FindSliderById(id)
	if err != nil {
		return slider, err
	}

	return slider, nil
}

func (s *service) DelSliderService(id int) (bool, error) {
	slider, err := s.repository.FindSliderById(id)
	if err != nil {
		return false, err
	}

	_, err = s.repository.DelSlider(id, slider)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *service) UpdateSliderByPostService(idslider int, idproduct int) (Slider, error) {
	slider, err := s.repository.FindSliderById(idslider)
	if err != nil {
		return slider, err
	}

	slider.ProductID = idproduct
	update, err := s.repository.UpdateSlider(slider)
	if err != nil {
		return update, err
	}

	return update, nil
}

// ============ PRODUCT
func (s *service) CreateProductByName(input CreateProductByName) (Product, error) {
	product := Product{}
	product.Name = input.Name
	product.Slug = input.Slug
	product.Active = 0

	create, err := s.repository.Save(product)
	if err != nil {
		return create, err
	}

	return create, nil
}

func (s *service) FindAllProductService() ([]Product, error) {
	products, err := s.repository.FindAll()
	if err != nil {
		return products, err
	}
	return products, nil
}

func (s *service) FindProductById(id int) (Product, error) {
	product, err := s.repository.FindById(id)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *service) FindProductBySlug(slug string) (bool, error) {
	product, err := s.repository.FindBySlug(slug)
	if err != nil {
		return false, err
	}

	if product.ID == 0 {
		return true, err
	}

	return false, nil
}

func (s *service) CreateProductService(input CreateProductInput) (Product, error) {
	product := Product{}
	product.Name = input.Name
	product.Slug = input.Slug
	product.Bahan = input.Bahan
	product.Price = input.Price
	product.Stock = input.Stock
	product.Description = input.Description
	product.Active = input.Active

	create, err := s.repository.Save(product)
	if err != nil {
		return create, err
	}

	return create, nil
}

func (s *service) UpdateProductService(input UpdateProductInput) (Product, error) {
	product, err := s.repository.FindById(input.ID)
	if err != nil {
		return product, err
	}

	product.Name = input.Name
	product.Slug = input.Slug
	product.Bahan = input.Bahan
	product.Price = input.Price
	product.Stock = input.Stock
	product.Description = input.Description
	product.Active = input.Active

	update, err := s.repository.Update(product)
	if err != nil {
		return update, err
	}

	return update, nil
}

func (s *service) DelProductService(id int) (bool, error) {
	product, err := s.repository.FindById(id)
	if err != nil {
		return false, err
	}

	_, err = s.repository.DelProduct(id, product)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *service) UpdateProductByActiveService(id int, input UpdateProductByActiveInput) (Product, error) {
	product, err := s.repository.FindById(id)
	if err != nil {
		return product, err
	}

	product.Active = input.Active
	update, err := s.repository.Update(product)
	if err != nil {
		return update, err
	}
	return update, nil
}

// ========== DISCOUNT
func (s *service) FindAllDiscountService() ([]Discount, error) {
	discounts, err := s.repository.FindAllDiscount()
	if err != nil {
		return discounts, err
	}
	return discounts, nil
}

func (s *service) CreateDiscountService(input CreateDiscountInput) (Discount, error) {
	discount := Discount{}
	discount.Name = input.Name
	discount.Slug = input.Slug
	discount.Persentase = input.Persentase
	discount.Price = input.Price
	discount.Active = input.Active
	discount.ProductID = input.ProductID
	// discount.StartDate = input.StartDate
	// discount.EndDate = input.EndDate

	create, err := s.repository.CreateDiscount(discount)
	if err != nil {
		return create, err
	}

	return create, nil
}

func (s *service) UpdateDiscountService(id int, input CreateDiscountInput) (Discount, error) {
	discount, err := s.repository.FindDiscountById(id)
	if err != nil {
		return discount, err
	}

	discount.Name = input.Name
	discount.Slug = input.Slug
	discount.Persentase = input.Persentase
	discount.Price = input.Price
	discount.Active = input.Active
	discount.ProductID = input.ProductID
	// discount.StartDate = input.StartDate
	// discount.EndDate = input.EndDate

	update, err := s.repository.UpdateDiscount(discount)
	if err != nil {
		return update, err
	}

	return update, nil
}

func (s *service) UpdateDiscountByActiveService(id int, input UpdateDiscountByActiveInput) (Discount, error) {
	discount, err := s.repository.FindDiscountById(id)
	if err != nil {
		return discount, err
	}

	discount.Active = input.Active

	update, err := s.repository.UpdateDiscount(discount)
	if err != nil {
		return update, err
	}

	return update, nil
}

func (s *service) DelDiscountService(id int) (bool, error) {
	discount, err := s.repository.DelDiscount(id)
	if err != nil {
		return false, err
	}

	if discount != true {
		return false, err
	}

	return true, err
}
