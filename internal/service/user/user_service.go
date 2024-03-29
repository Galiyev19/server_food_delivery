package user

import (
	"fmt"
	"food_delivery/internal/models"
	"food_delivery/internal/repository/user"

	"github.com/google/uuid"
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
	UpdateUser(username, email string) error
	GetUserByEmail(email string) (models.User, error)
}

func (u *UserService) CreateUser(user models.User) error {
	id := uuid.New()
	userModel := models.User{
		ID:       id.String(),
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}

	user, _ = u.UserRepository.GetUserByEmail(user.Email)
	if user.Email == user.Email {
		return fmt.Errorf("User this email already exist")
	}

	if err := u.UserRepository.CreateUser(userModel); err != nil {
		return fmt.Errorf("u.UserRepository.CreateUser: %v", err)
	}
	return nil
}

func (u *UserService) GetUserByEmail(email string) (models.User, error) {
	return u.UserRepository.GetUserByEmail(email)
}

func (u *UserService) UpdateUser(username, email string) error {
	user, err := u.UserRepository.GetUserByEmail(username)
	if err != nil {
		return fmt.Errorf("User current email not exist %s - ", username)
	}

	err = u.UserRepository.UpdateUser(user.UserName, user.Email)
	if err != nil {
		return fmt.Errorf("NOT UPDATE")
	}

	return nil
}
