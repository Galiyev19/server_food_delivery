package handlers

import (
	"fmt"
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

	adminResponse, err := h.service.Admin.CreateAdmin(adminModel)
	if err != nil {
		h.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	err = h.writeJson(w, http.StatusOK, envelope{"admin": adminResponse}, nil)
	if err != nil {
		h.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}
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
	fmt.Println(admin)

	err = h.writeJson(w, http.StatusOK, envelope{"data": admin}, nil)
}

func (h *Handler) IdentityMe(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("authorization")  // token from header
	tokenSplit := strings.Split(token, " ") // split token
	data, err := h.service.Admin.IdentityMe(tokenSplit[1])
	if err != nil {
		h.errorResponse(w, r, 401, err.Error())
		return
	}

	err = h.writeJson(w, http.StatusOK, envelope{"data": data}, nil)
}

func (h *Handler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	// check token
	token := r.Header.Get("authorization")
	_, err := h.service.Admin.IdentityMe(token)
	if err != nil {
		h.errorResponse(w, r, 401, err.Error())
		return
	}

	// input data
	var inp struct {
		Email       string `json:"email"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	// read data json data
	err = h.readJson(w, r, &inp)
	if err != nil {
		h.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// check admin is exits in db
	_, err = h.service.Admin.GetAdminInfo(inp.Email)
	if err != nil {
		h.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Admin.ChangePassword(inp.Email, inp.OldPassword, inp.NewPassword)
	if err != nil {
		h.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	err = h.writeJson(w, http.StatusOK, envelope{"message": "password was changed"}, nil)
}
