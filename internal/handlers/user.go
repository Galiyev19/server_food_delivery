package handlers

import (
	"food_delivery/internal/models"
	"food_delivery/internal/validator"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID       string `json:"id"`
		UserName string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	user := models.User{
		ID:       input.ID,
		UserName: input.UserName,
		Email:    input.Email,
		Password: input.Password,
	}

	err := h.readJson(w, r, &user)
	if err != nil {
		h.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	v := validator.New()

	if models.ValidateUser(v, &user); !v.Valid() {
		h.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = h.service.User.CreateUser(user)
	if err != nil {
		h.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	userResponse := models.UserResponse{
		UserName: user.UserName,
		Email:    user.Email,
	}

	err = h.writeJson(w, http.StatusOK, envelope{"user": userResponse}, nil)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	username, err := h.readStrParam(r)
	if err != nil {
		h.notFoundResponse(w, r)
		return
	}

	
}
