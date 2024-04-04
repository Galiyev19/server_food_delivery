package handlers

import (
	"food_delivery/internal/models"
	"food_delivery/internal/validator"
	"net/http"
	"strings"
)

func (h *Handler) CreateAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		h.methodNotAllowed(w, r)
		return
	}

	var inp struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	adminModel := models.Admin{
		Email:    inp.Email,
		Password: inp.Password,
	}

	err := h.readJson(w, r, &adminModel)
	if err != nil {
		h.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	v := validator.New()

	if models.ValidAdmin(v, &adminModel); !v.Valid() {
		h.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = h.service.Admin.CreateAdmin(adminModel)
	if err != nil {
		h.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	adminResponse := models.AdminResponse{
		Email: adminModel.Email,
	}

	err = h.writeJson(w, http.StatusOK, envelope{"admin": adminResponse}, nil)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.methodNotAllowed(w, r)
		return
	}

	var inp struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	adminModel := models.Admin{
		Email:    inp.Email,
		Password: inp.Password,
	}

	err := h.readJson(w, r, &adminModel)
	if err != nil {
		h.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	admin, err := h.service.Admin.GetAdminByEmail(adminModel.Email, adminModel.Password)
	if err != nil {
		h.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	err = h.writeJson(w, http.StatusOK, envelope{"data": admin}, nil)
}

func (h *Handler) IdentityMe(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("authorization")
	parts := strings.Split(header, " ")

	data, err := h.service.Admin.IdentityMe(parts[1])
	if err != nil {
		h.errorResponse(w, r, 401, err.Error())
		return
	}

	res := models.Response{
		Email:     data.Email,
		ExpiresAt: data.ExpiresAt,
		IssuedAt:  data.IssuedAt,
	}

	err = h.writeJson(w, http.StatusOK, envelope{"data": res}, nil)
}
