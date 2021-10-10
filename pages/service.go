package pages

type Service interface {
	FindAll() ([]Page, error)
	FindById(id int) (Page, error)
	FindBySlug(slug string) (Page, error)
	CreatePage(input CreatePageInput) (Page, error)
	UpdatePage(input UpdatePageInput) (Page, error)
	DelPage(id int) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Page, error) {
	pages, err := s.repository.FindAll()
	if err != nil {
		return pages, err
	}
	return pages, nil
}

func (s *service) FindById(id int) (Page, error) {
	page, err := s.repository.FindById(id)
	if err != nil {
		return page, err
	}
	return page, nil
}

func (s *service) FindBySlug(slug string) (Page, error) {
	page, err := s.repository.FindBySlug(slug)
	if err != nil {
		return page, err
	}
	return page, nil
}

func (s *service) CreatePage(input CreatePageInput) (Page, error) {
	page := Page{}
	page.Name = input.Name
	page.Slug = input.Slug
	page.Description = input.Description

	page, err := s.repository.Save(page)
	if err != nil {
		return page, err
	}
	return page, nil
}

func (s *service) UpdatePage(input UpdatePageInput) (Page, error) {
	page, err := s.repository.FindById(input.ID)

	if err != nil {
		return page, err
	}

	page.Name = input.Name
	page.Slug = input.Slug
	page.Description = input.Description

	update, err := s.repository.Update(page)
	if err != nil {
		return update, err
	}
	return update, nil
}

func (s *service) DelPage(id int) (bool, error) {
	page, err := s.repository.FindById(id)
	if err != nil {
		return false, err
	}

	_, err = s.repository.Delete(id, page)
	if err != nil {
		return false, err
	}

	return true, nil
}
