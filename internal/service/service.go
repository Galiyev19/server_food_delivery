package service

import (
	"food_delivery/internal/repository"
	"food_delivery/internal/service/user"
)

type Service struct {
	User user.IUserService
}

func NewService(r *repository.Repositories) *Service {
	return &Service{
		User: user.NewUserService(&r.User),
	}
}
