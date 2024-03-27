package user

import (
	"fmt"
	"food_delivery/internal/models"
	"food_delivery/internal/repository/user"
)

type UserService struct {
	UserRepository user.IUserRepository
}

func NewUserService(userRepo user.IUserRepository) *UserService {
	return &UserService{
		UserRepository: userRepo,
	}
}

type IUserService interface {
	CreateUser(user models.User) error
	GetUserById() error
	UpdateUser(user models.User) error
}

func (u *UserService) CreateUser(user models.User) error {
	userModel := models.User{
		ID:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}

	fmt.Println(userModel)

	if err := u.UserRepository.CreateUser(userModel); err != nil {
		return fmt.Errorf("u.UserRepository.CreateUser: %v", err)
	}
	return nil
}

func (u *UserService) GetUserById() error {
	return nil
}

func (u *UserService) UpdateUser(user models.User) error {
	return nil
}
