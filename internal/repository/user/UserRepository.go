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
	GetUserByEmail(email string) (models.User, error)
	UpdateUser(username, email string) error
}

func (u *UserRepository) CreateUser(user models.User) error {
	stmt := `INSERT INTO users (id, username, email, password,created_at) 
		VALUES (?, ?, ?, ?, datetime('now', 'localtime'));`
	if _, err := u.db.Exec(stmt, user.ID, user.UserName, user.Email, user.Password); err != nil {
		return fmt.Errorf("u.db.Exec: %v", err)
	}
	return nil
}

func (u *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	stmt := `SELECT * FROM users WHERE email = ?`
	if err := u.db.QueryRow(stmt, email).Scan(&user.ID, &user.UserName, &user.Email, &user.Password); err != nil {
		return models.User{}, fmt.Errorf("NOT FIND USER")
	}
	return user, nil
}

func (u *UserRepository) UpdateUser(username, email string) error {
	stmt := `UPDATE users SET username = ?. email = ? WHERE username = ?`
	if _, err := u.db.Exec(stmt, username, email); err != nil {
		return fmt.Errorf("UPDATE USER %v", err)
	}
	return nil
}
