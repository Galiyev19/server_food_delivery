package models

import "time"

type Token struct {
	UserName string `json:"username"`
	Token    string `json:"token"`
}

type TokenResponse struct {
	ExpiresAt time.Time `json:"exp"`
	IssuedAt  time.Time `json:"iat"`
}
