package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) Routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", h.healthCheck)

	// User
	router.HandlerFunc(http.MethodPost, "/v1/sign-up", h.CreateUser)
	router.HandlerFunc(http.MethodPatch, "/v1/user/:username", h.UpdateUser)

	// Admin
	router.HandlerFunc(http.MethodPost, "/v1/admin/sign-up", h.CreateAdmin)

	return router
}
