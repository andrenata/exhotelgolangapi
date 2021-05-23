package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(id int, fileLocation string) (User, error)
	GetUserbyId(id int) (User, error)
	ServiceChangeName(id int, input ChangeNameInput) (User, error)
	ServiceCheckPin(id int, input CheckPin) (bool, error)
	ServiceChangePin(id int, input ChangePin) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.TypeVerified = 0
	user.IsVerified = 0
	user.IsActive = 0
	user.PhoneNumber = input.PhoneNumber
	Password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(Password)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil

}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, nil
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

// mapping struct input ke struct user
// simpan struct user ke responsitory

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(id int, fileLocation string) (User, error) {
	// dapatkan user by ID
	// update attribute avatar file name
	// simpan perubahan avatar file name

	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	user.PicturePath = fileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserbyId(id int) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	return user, nil
}

func (s *service) ServiceChangeName(id int, input ChangeNameInput) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	user.Name = input.Name
	updatedName, err := s.repository.Update(user)
	if err != nil {
		return updatedName, err
	}

	return updatedName, nil

}

func (s *service) ServiceChangePin(id int, input ChangePin) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	Pin, err := bcrypt.GenerateFromPassword([]byte(input.Pin), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Pin = string(Pin)
	updatePin, err := s.repository.Update(user)
	if err != nil {
		return updatePin, err
	}

	return updatePin, nil

}

func (s *service) ServiceCheckPin(id int, input CheckPin) (bool, error) {

	user, err := s.repository.FindById(id)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return false, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Pin), []byte(input.Pin))
	if err != nil {
		return false, nil
	}

	return true, nil

}
