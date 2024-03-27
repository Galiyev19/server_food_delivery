package handlers

import (
	"food_delivery/internal/models"
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

	err = h.service.User.CreateUser(user)
	if err != nil {
		h.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	err = h.writeJson(w, http.StatusOK, envelope{"user": user}, nil)
}
