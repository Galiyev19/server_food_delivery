package models

import "time"

type Token struct {
	UserName string `json:"username"`
	Token    string `json:"token"`
}

type Response struct {
	Email     string    `json:"sub"`
	ExpiresAt time.Time `json:"exp"`
	IssuedAt  time.Time `json:"iat"`
}
