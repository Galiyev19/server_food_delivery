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
	CreateAdmin(admin models.Admin) (models.AdminResponse, error)
	GetAdminByEmail(email, password string) (models.Token, error)
	IdentityMe(token string) (models.AdminResponse, error)
	GetAdminInfo(email string) (models.AdminResponse, error)
	ChangePassword(email, oldPassword, newPassword string) error
}

func (a *AdminService) CreateAdmin(admin models.Admin) (models.AdminResponse, error) {
	id := helpers.GenerateId()                            // generate id
	hashPass, err := helpers.HashPassword(admin.Password) // hashed password
	if err != nil {
		return models.AdminResponse{}, fmt.Errorf("Create admin hash password")
	}

	adminModel := models.Admin{
		ID:       id.String(),
		Email:    admin.Email,
		Password: hashPass,
	}

	if err := a.AdminRepository.CreateAdmin(adminModel); err != nil {
		return models.AdminResponse{}, fmt.Errorf("Service create admin - %v", err)
	}

	token, err := helpers.GenerateToken(admin.Email)
	if err != nil {
		return models.AdminResponse{}, err
	}

	data, err := a.IdentityMe(token)
	if err != nil {
		return models.AdminResponse{}, err
	}

	return data, nil
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

	claims := models.Token{
		UserName: admin.Email,
		Token:    tokenString,
	}
	return claims, nil
}

func (a *AdminService) IdentityMe(token string) (models.AdminResponse, error) {
	claims, err := helpers.ParseToken(token)
	if err != nil {
		return models.AdminResponse{}, fmt.Errorf("401 unauthorized")
	}

	dateExp, err := helpers.ConvertToDate(claims["exp"].(float64))
	if err != nil {
		return models.AdminResponse{}, fmt.Errorf("Identity me %v", err)
	}

	dateIssued, err := helpers.ConvertToDate(claims["iat"].(float64))
	if err != nil {
		return models.AdminResponse{}, fmt.Errorf("Identity me %v", err)
	}

	tokenResponse := models.TokenResponse{
		ExpiresAt: dateExp,
		IssuedAt:  dateIssued,
	}

	data := models.AdminResponse{
		Email: claims["sub"].(string),
		Token: tokenResponse,
	}

	return data, nil
}

func (a *AdminService) GetAdminInfo(email string) (models.AdminResponse, error) {
	data, err := a.AdminRepository.GetAdminByEmail(email)
	if err != nil {
		return models.AdminResponse{}, fmt.Errorf("Admin not found")
	}

	res := models.AdminResponse{
		Email: data.Email,
	}
	return res, nil
}

func (a *AdminService) ChangePassword(email, oldPassword, newPassword string) error {
	// get info current admin
	admin, err := a.AdminRepository.GetAdminByEmail(email)
	if err != nil {
		return fmt.Errorf("service error not find admin %v", err)
	}

	// compare current password and password from input
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(oldPassword))
	if err != nil {
		return fmt.Errorf("wrong password")
	}
	// changed password
	err = a.AdminRepository.ChangePassword(email, newPassword)
	if err != nil {
		return err
	}
	return nil
}
