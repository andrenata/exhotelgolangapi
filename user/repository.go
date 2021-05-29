package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(id int) (User, error)
	Update(user User) (User, error)
	FindByPhone(PhoneNumber string) (User, error)
	// CheckPin(id int, pin string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *repository) FindByPhone(PhoneNumber string) (User, error) {
	var user User
	err := r.db.Where("phone_number = ?", PhoneNumber).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *repository) FindById(id int) (User, error) {
	var user User

	err := r.db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// func (r *repository) CheckPin(id int, pin string) (User, error) {
// 	var user User

// 	err := r.db.Where("id = ? AND pin = ?", id, pin).Find(&user).Error
// 	if err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }
