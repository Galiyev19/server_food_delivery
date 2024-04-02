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
	Token string
}

func ValidAdmin(v *validator.Validator, a *Admin) {
	ValidateEmail(v, a.Email)
	ValidatePaswword(v, a.Password)
}
