package services

import (
	"sudut_literat/app/request"
	"sudut_literat/app/response"
	"sudut_literat/helper"
	"sudut_literat/models"
	"sudut_literat/repositories"

	"github.com/go-playground/validator/v10"
)

type UserService interface {
	GetAll() []response.UserResponse
	GetDetail(id int) response.UserResponse
	CreateUser(user request.UserRequest)
}

type userService struct {
	UserRepository repositories.UserRepository
	Validate       *validator.Validate
}

// CreateUser implements UserService
func (u *userService) CreateUser(user request.UserRequest) {
	err := u.Validate.Struct(user)
	helper.Error(err)

	userModel := models.User{
		Name:  user.Name,
		Email: user.Email,
		Password: user.Password,
	}

	u.UserRepository.Create(userModel)
}

// GetAll implements UserService
func (u *userService) GetAll() []response.UserResponse {
	res := u.UserRepository.List()

	users := []response.UserResponse{}
	for _, value := range res {
		user := response.UserResponse{
			ID:        int(value.ID),
			Name:      value.Name,
			Email:     value.Email,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		users = append(users, user)
	}
	return users
}

// GetDetail implements UserService
func (u *userService) GetDetail(id int) response.UserResponse {
	res, err:= u.UserRepository.Detail(id)
	helper.Error(err)
	user := response.UserResponse{
		ID:        int(res.ID),
		Name:      res.Name,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}
	return user
}

func NewUserService(userRepository repositories.UserRepository, validate *validator.Validate) UserService {
	return &userService{
		UserRepository: userRepository,
		Validate: 	 validate,
	}
}
