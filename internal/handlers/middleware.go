package handlers

import (
	"food_delivery/internal/service/helpers"
	"net/http"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		if authToken == "" {
			h.errorResponse(w, r, http.StatusUnauthorized, "Not authorized")
			return
		}

		_, err := helpers.ParseToken(authToken)
		if err != nil {
			h.errorResponse(w, r, http.StatusUnauthorized, err.Error())
			return
		}

		next(w, r)
	}
}
