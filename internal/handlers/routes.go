package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) Routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", h.healthCheck)
	router.HandlerFunc(http.MethodPost, "/v1/user", h.CreateUser)

	return router
}
