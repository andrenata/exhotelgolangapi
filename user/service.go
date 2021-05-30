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
	ServiceChangePhoneNumber(id int, input InputChangeNumber) (User, error)
	ChangeEmailService(id int, input ChangeEmailInput) (User, error)
	IsPhoneAvailable(input CheckPhoneInput) (bool, error)
	ServiceCheckPinTemporary(id int, input CheckPin) (bool, error)
	ServiceChangePinTemporary(id int, input ChangePinTemporary) (User, error)
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
	user.Balance = 0
	user.BalanceTemporary = 0
	user.PhoneNumber = input.PhoneNumber
	Password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(Password)

	checkPhone, err := s.repository.FindByPhone(input.PhoneNumber)
	if err != nil {
		return checkPhone, err
	}

	if checkPhone.ID != 0 {
		return checkPhone, errors.New("Phone number has been registered")
	}

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

func (s *service) IsPhoneAvailable(input CheckPhoneInput) (bool, error) {
	phone := input.PhoneNumber

	user, err := s.repository.FindByPhone(phone)
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
	newpin := input.Pin
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, nil
	}
	if user.ID == 0 {
		return user, errors.New("PIN is different")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PinTemporary), []byte(newpin))
	if err != nil {
		return user, err
	}

	pin, err := bcrypt.GenerateFromPassword([]byte(input.Pin), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Pin = string(pin)
	user.PinTemporary = ""
	updatePin, err := s.repository.Update(user)
	if err != nil {
		return updatePin, err
	}

	return updatePin, nil

}

func (s *service) ServiceChangePinTemporary(id int, input ChangePinTemporary) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	Pin, err := bcrypt.GenerateFromPassword([]byte(input.PinTemporary), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PinTemporary = string(Pin)
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

func (s *service) ServiceCheckPinTemporary(id int, input CheckPin) (bool, error) {

	user, err := s.repository.FindById(id)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return false, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PinTemporary), []byte(input.Pin))
	if err != nil {
		return false, nil
	}

	return true, nil

}

func (s *service) ServiceChangePhoneNumber(id int, input InputChangeNumber) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	user.PhoneNumber = input.PhoneNumber
	changedPhone, err := s.repository.Update(user)
	if err != nil {
		return changedPhone, err
	}
	return changedPhone, nil
}

func (s *service) ChangeEmailService(id int, input ChangeEmailInput) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	user.Email = input.Email
	changedEmail, err := s.repository.Update(user)
	if err != nil {
		return changedEmail, err
	}
	return changedEmail, nil
}
