package service

import (
	"food_delivery/internal/repository"
	"food_delivery/internal/service/admin"
	"food_delivery/internal/service/products"
	"food_delivery/internal/service/user"
)

type Service struct {
	User    user.IUserService
	Admin   admin.IAdminService
	Product products.IProductsService
}

func NewService(r *repository.Repositories) *Service {
	return &Service{
		User:    user.NewUserService(&r.User),
		Admin:   admin.NewAdminService(&r.Admin),
		Product: products.NewProductService(&r.Product),
	}
}
