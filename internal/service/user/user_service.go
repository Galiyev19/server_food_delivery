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

// Create new user and save in db
func (u *UserService) CreateUser(user models.User) error {
	id := uuid.New()                                 // generate uuid
	hashPassword, err := HashPassword(user.Password) // hash password
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	userModel := models.User{
		ID:       id.String(),
		UserName: user.UserName,
		Email:    user.Email,
		Password: hashPassword,
	}

	userFromDB, _ := u.UserRepository.GetUserByEmail(user.Email) // check this email is exist

	if userFromDB.Email == user.Email {
		return fmt.Errorf("User this email already exist")
	}

	// create new user
	if err := u.UserRepository.CreateUser(userModel); err != nil {
		return fmt.Errorf("Service error create user: %v", err)
	}
	return nil
}

// get user by email
func (u *UserService) GetUserByEmail(email string) (models.User, error) {
	return u.UserRepository.GetUserByEmail(email)
}

// update user info
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
