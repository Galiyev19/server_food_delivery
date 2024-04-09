package handlers

import (
	"food_delivery/internal/models"
	"food_delivery/internal/validator"
	"net/http"
)

func (h *Handler) InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		h.methodNotAllowed(w, r)
		return
	}

	var rating struct {
		Rate  float64 `json:"rate"`
		Count int64   `json:"count"`
	}
	var input struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Category    string  `json:"category"`
		Price       float64 `json:"price"`
		Image       string  `json:"image"`
	}

	products := models.Products{
		Title:       input.Title,
		Description: input.Description,
		Category:    input.Category,
		Price:       input.Price,
		Image:       input.Image,
		Rating:      rating,
	}

	err := h.readJson(w, r, &products)
	if err != nil {
		h.errorResponse(w, r, http.StatusBadGateway, err.Error())
		return
	}

	v := validator.New()
	if models.ValidatorProduct(v, &products); !v.Valid() {
		h.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = h.writeJson(w, http.StatusOK, envelope{"message": "inserted"}, nil)
}