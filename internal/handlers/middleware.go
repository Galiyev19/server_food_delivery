package handlers

import (
	"food_delivery/internal/service/helpers"
	"net/http"
)

const (
	authorizationHeader = "Authorization"
)

// AuthMiddleware принимает обычную функцию
func (h *Handler) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Здесь вы можете добавить логику проверки авторизации
		// Например, проверить наличие токена аутентификации в заголовке запроса
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

		// Если пользователь авторизован, переходим к следующему обработчику
		next(w, r)
	}
}
