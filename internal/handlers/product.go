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

	var input struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Category    string  `json:"category"`
		Price       float64 `json:"price"`
		Image       string  `json:"image"`
		Rating      struct {
			Rate  float64 `json:"rate"`
			Count int64   `json:"count"`
		} `json:"rating"`
	}

	ratingModel := models.Rating{
		Rate:  input.Rating.Rate,
		Count: input.Rating.Count,
	}

	products := models.Products{
		Title:       input.Title,
		Description: input.Description,
		Category:    input.Category,
		Price:       input.Price,
		Image:       input.Image,
		Rating:      ratingModel,
	}

	// parse json
	err := h.readJson(w, r, &products)
	if err != nil {
		h.errorResponse(w, r, http.StatusBadGateway, err.Error())
		return
	}

	// validator for data
	v := validator.New()
	if models.ValidatorProduct(v, &products); !v.Valid() {
		h.failedValidationResponse(w, r, v.Errors)
		return
	}

	// insert data
	err = h.service.Product.InsertProduct(products)
	if err != nil {
		h.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// return json
	err = h.writeJson(w, http.StatusOK, envelope{"message": products}, nil)
	if err != nil {
		h.errorResponse(w, r, http.StatusBadGateway, err.Error())
		return
	}
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := h.readIDParam(r)
	if err != nil {
		h.notFoundResponse(w, r)
		return
	}

	_, err = h.service.Product.GetProductById(id)
	if err != nil {
		h.notFoundResponse(w, r)
		return
	}

	var input struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Category    string  `json:"category"`
		Price       float64 `json:"price"`
		Image       string  `json:"image"`
		Rating      struct {
			Rate  float64 `json:"rate"`
			Count int64   `json:"count"`
		} `json:"rating"`
	}

	ratingModel := models.Rating{
		Rate:  input.Rating.Rate,
		Count: input.Rating.Count,
	}

	productModel := models.Products{
		Title:       input.Title,
		Description: input.Description,
		Category:    input.Category,
		Price:       input.Price,
		Image:       input.Image,
		Rating:      ratingModel,
	}

	err = h.readJson(w, r, &productModel)
	if err != nil {
		h.errorResponse(w, r, http.StatusBadGateway, err.Error())
		return
	}

	err = h.service.Product.UpdateProduct(id, productModel)
	if err != nil {
		h.errorResponse(w, r, http.StatusBadGateway, err.Error())
	}

	h.writeJson(w, http.StatusOK, envelope{"response": "changed"}, nil)
}
