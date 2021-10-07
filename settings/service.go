package settings

type Service interface {
	FindById(id int) (Setting, error)
	UpdateSetting(input UpdateSettingInput) (Setting, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindById(id int) (Setting, error) {
	setting, err := s.repository.FindById(id)
	if err != nil {
		return setting, err
	}
	return setting, nil
}

func (s *service) UpdateSetting(input UpdateSettingInput) (Setting, error) {
	setting, err := s.repository.FindById(input.ID)

	if err != nil {
		return setting, err
	}

	setting.Title = input.Title
	setting.Keyword = input.Keyword
	setting.Facebook = input.Facebook
	setting.Instagram = input.Instagram
	setting.Maps = input.Maps
	setting.Address = input.Address
	setting.Phone = input.Phone
	setting.Description = input.Description

	update, err := s.repository.Update(setting)
	if err != nil {
		return update, err
	}
	return update, nil
}
