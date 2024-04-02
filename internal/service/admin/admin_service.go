package admin

import (
	"fmt"
	"food_delivery/internal/models"
	"food_delivery/internal/repository/admin"

	"github.com/google/uuid"
)

type AdminService struct {
	AdminRepository admin.IAdminRepository
}

func NewAdminService(adminRepo admin.IAdminRepository) *AdminService {
	return &AdminService{
		AdminRepository: adminRepo,
	}
}

type IAdminService interface {
	CreateAdmin(admin models.Admin) error
	GetAdminByEmail(email string) (models.Admin, error)
}

func (a *AdminService) CreateAdmin(admin models.Admin) error {
	id := uuid.New()

	adminModel := models.Admin{
		ID:       id.String(),
		Email:    admin.Email,
		Password: admin.Password,
	}

	if err := a.AdminRepository.CreateAdmin(adminModel); err != nil {
		return fmt.Errorf("Service create admin - %v", err)
	}
	return nil
}

func (a *AdminService) GetAdminByEmail(email string) (models.Admin, error) {
	var admin models.Admin

	return admin, nil
}
