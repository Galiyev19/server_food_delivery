package admin

import (
	"fmt"
	"food_delivery/internal/models"
	"food_delivery/internal/repository/admin"
	"food_delivery/internal/service/helpers"

	"golang.org/x/crypto/bcrypt"
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
	GetAdminByEmail(email, password string) (models.Token, error)
	IdentityMe(token string) (models.Response, error)
}

func (a *AdminService) CreateAdmin(admin models.Admin) error {
	id := helpers.GenerateId()                            // generate id
	hashPass, err := helpers.HashPassword(admin.Password) // hashed password
	if err != nil {
		return fmt.Errorf("Create admin hash password")
	}

	adminModel := models.Admin{
		ID:       id.String(),
		Email:    admin.Email,
		Password: hashPass,
	}

	if err := a.AdminRepository.CreateAdmin(adminModel); err != nil {
		return fmt.Errorf("Service create admin - %v", err)
	}
	return nil
}

func (a *AdminService) GetAdminByEmail(email, password string) (models.Token, error) {
	admin, err := a.AdminRepository.GetAdminByEmail(email)
	if err != nil {
		return models.Token{}, fmt.Errorf("service error not find admin %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		return models.Token{}, fmt.Errorf("wrong password")
	}

	tokenString, err := helpers.GenerateToken(admin.Email)
	if err != nil {
		return models.Token{}, err
	}

	fmt.Println(tokenString)
	claims := models.Token{
		UserName: admin.Email,
		Token:    tokenString,
	}
	return claims, nil
}

func (a *AdminService) IdentityMe(token string) (models.Response, error) {
	claims, err := helpers.ParseToken(token)
	if err != nil {
		return models.Response{}, fmt.Errorf("Identity me %v", err)
	}

	dateExp, err := helpers.ConvertToDate(claims["exp"].(float64))
	if err != nil {
		return models.Response{}, fmt.Errorf("Identity me %v", err)
	}
	dateIssued, err := helpers.ConvertToDate(claims["iat"].(float64))
	if err != nil {
		return models.Response{}, fmt.Errorf("Identity me %v", err)
	}
	data := models.Response{
		Email:     claims["sub"].(string),
		ExpiresAt: dateExp,
		IssuedAt:  dateIssued,
	}

	return data, nil
}
