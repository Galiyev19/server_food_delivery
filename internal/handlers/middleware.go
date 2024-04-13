package handlers

import (
	"food_delivery/internal/service/helpers"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		if len(bearerToken) == 0 {
			h.errorResponse(w, r, http.StatusUnauthorized, "Not authorized")
			return
		}

		authToken := strings.Split(bearerToken, " ")
		if authToken[1] == "" {
			h.errorResponse(w, r, http.StatusUnauthorized, "Not authorized")
			return
		}

		_, err := helpers.ParseToken(authToken[1])
		if err != nil {
			h.errorResponse(w, r, http.StatusUnauthorized, err.Error())
			return
		}

		next(w, r)
	}
}
