package category

import "strings"

type Service interface {
	RegisterCategory(input RegisterCategoryInput) (Category, error)
	UpdateCategory(input UpdateCategoryInput) (Category, error)
	GetCategoryById(input GetCategoryDetailInput) (Category, error)
	GetCategoryBySlug(input GetCategorySlugInput) (Category, error)
	DeleteCategory(input DeleteCategoryInput) (bool, error)
	FindAllCategory() ([]Category, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCategoryById(input GetCategoryDetailInput) (Category, error) {
	category, err := s.repository.FindById(input.ID)

	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *service) FindAllCategory() ([]Category, error) {
	categories, err := s.repository.FindAll()

	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (s *service) GetCategoryBySlug(input GetCategorySlugInput) (Category, error) {
	category, err := s.repository.FindBySlug(input.Slug)

	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *service) RegisterCategory(input RegisterCategoryInput) (Category, error) {
	category := Category{}
	category.Name = input.Name
	category.Slug = strings.ReplaceAll(strings.ToLower(input.Name), " ", "-")

	find, err := s.repository.FindBySlug(category.Slug)
	if err != nil {
		return find, err
	}

	newCategory, err := s.repository.Save(category)
	if err != nil {
		return newCategory, err
	}

	return newCategory, nil

}

func (s *service) UpdateCategory(input UpdateCategoryInput) (Category, error) {
	category, err := s.repository.FindById(input.Id)
	if err != nil {
		return category, err
	}

	category.Name = input.Name
	update, err := s.repository.Update(category)
	if err != nil {
		return update, err
	}

	return update, nil

}

func (s *service) DeleteCategory(input DeleteCategoryInput) (bool, error) {
	category, err := s.repository.FindById(input.ID)
	if err != nil {
		return false, err
	}

	_, err = s.repository.Delete(input.ID, category)
	if err != nil {
		return false, err
	}
	return true, nil
}
