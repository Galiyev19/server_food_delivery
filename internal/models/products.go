package models

import (
	"food_delivery/internal/validator"
	"time"
)

type Rating struct {
	Rate  float64 `json:"rate"`
	Count int64   `json:"count"`
}

type Products struct {
	ID          int64     `json:"id"`
	CreadAt     time.Time `json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Price       int64     `json:"price"`
	Image       string    `json:"image"`
	Rating      Rating    `json:"rating"`
	Version     int64     `json:"version"`
}

func ValidatorProduct(v *validator.Validator, p *Products) {
	v.Check(p.Title != "", "title", "must be provided")
	v.Check(len(p.Title) <= 500, "title", "must be not more than 500 bytes long")

	v.Check(p.Description != "", "description", "must be provided")
	v.Check(len(p.Description) <= 500, "description", "must be not more than 500 bytes long")

	v.Check(p.Category != "", "category", "must be provided")
	v.Check(len(p.Category) <= 500, "category", "must be not more than 500 bytes long")

	v.Check(p.Price > 0, "price", "must ber positive integer")

	v.Check(p.Image != "", "image", "must be provided")
	v.Check(len(p.Image) <= 500, "image", "must be not more than 500 bytes long")

	v.Check(p.Rating.Rate > 0, "rate", "must be positive integer")
	v.Check(p.Rating.Count > 0, "count", "must be positive integer")
}
