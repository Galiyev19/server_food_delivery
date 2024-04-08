package models

import (
	"food_delivery/internal/validator"
	"time"
)

type Admin struct {
	ID        string
	Email     string
	Password  string
	CreatedAt time.Time
}

type AdminResponse struct {
	Email string
	Token TokenResponse
}

func ValidAdmin(v *validator.Validator, a *Admin) {
	ValidateEmail(v, a.Email)
	ValidatePaswword(v, a.Password)
}
