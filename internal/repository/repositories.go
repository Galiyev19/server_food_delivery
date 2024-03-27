package repository

import (
	"database/sql"
	"food_delivery/internal/repository/user"
)

type Repositories struct {
	User user.UserRepository
}

func NewRepository(db *sql.DB) *Repositories {
	return &Repositories{
		User: *user.NewUserRepository(db),
	}
}
