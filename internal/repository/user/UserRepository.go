package user

import (
	"database/sql"
	"fmt"
	"food_delivery/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type IUserRepository interface {
	CreateUser(user models.User) error
	GetUserById() error
	UpdateUser(user models.User) error
}

func (u *UserRepository) CreateUser(user models.User) error {
	stmt := `INSERT INTO users (id, username, email, password) 
		VALUES (?, ?, ?, ?);`
	if _, err := u.db.Exec(stmt, user.ID, user.UserName, user.Email, user.Password); err != nil {
		return fmt.Errorf("u.db.Exec: %v", err)
	}
	return nil
}

func (u *UserRepository) GetUserById() error {
	return nil
}

func (u *UserRepository) UpdateUser(user models.User) error {
	return nil
}
