package repositories

import (
	"errors"
	"sudut_literat/helper"
	"sudut_literat/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	List() []models.User
	Detail(id int) (users models.User, err error)
	// DetailByEmail(email string)
	Create(user models.User)
	// Update(id int, user request.UpdateUserRequest)
	// Delete(id int)
}

type userRepository struct {
	db *gorm.DB
}

// Create implements UserRepository
func (u *userRepository) Create(user models.User) {
	res := u.db.Create(&user)
	helper.Error(res.Error)
}

// Detail implements UserRepository
func (u *userRepository) Detail(id int) (users models.User, err error) {
	var user models.User
	res := u.db.First(&user, id)
	if res != nil {
		return user, nil
	} else{
		return user, errors.New("user not found")
	}
}



// List implements UserRepository
func (u *userRepository) List() []models.User {
	var users []models.User
	res := u.db.Find(&users)
	helper.Error(res.Error)
	return users
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
