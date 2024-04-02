package models

import (
	"food_delivery/internal/validator"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserResponse struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")                                      // not be empty
	v.Check(v.Matches(email, validator.EmailRX), "email", "must be a valid email address") // check valid email
}

func ValidatePaswword(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")                 // not be empty
	v.Check(len(password) >= 8, "password", "must be at least 8 character") // min len password
}

func ValidateUser(v *validator.Validator, u *User) {
	v.Check(u.UserName != "", "username", "must be provided")                               // check empty username
	v.Check(len(u.UserName) > 3, "username", "username must be at least 3 characters long") // username not be less 3
	v.Check(len(u.UserName) < 50, "username", "must not exceed 50 characters")              // username not be more than 50

	ValidateEmail(v, u.Email)       // check email
	ValidatePaswword(v, u.Password) // check password
}
