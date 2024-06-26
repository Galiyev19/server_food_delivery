package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) Routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", h.healthCheck)

	// User
	router.HandlerFunc(http.MethodPost, "/v1/sign-up", h.CreateUser)
	router.HandlerFunc(http.MethodPatch, "/v1/user/:username", h.UpdateUser)

	// Admin
	router.HandlerFunc(http.MethodPost, "/v1/admin/sign-up", h.CreateAdmin)
	router.HandlerFunc(http.MethodPost, "/v1/admin/sign-in", h.Login)
	router.HandlerFunc(http.MethodGet, "/v1/admin/identity/me", h.IdentityMe)

	router.HandlerFunc(http.MethodPatch, "/v1/update/admin", h.AuthMiddleware(h.ChangePassword))

	// Product
	router.HandlerFunc(http.MethodPost, "/v1/product", h.AuthMiddleware(h.InsertProduct))
	router.HandlerFunc(http.MethodPatch, "/v1/product/:id", h.AuthMiddleware(h.UpdateProduct))
	router.HandlerFunc(http.MethodDelete, "/v1/product/:id", h.AuthMiddleware(h.DeleteProduct))
	router.HandlerFunc(http.MethodGet, "/v1/product/:id", h.AuthMiddleware(h.GetProduct))

	corsHandler := corsHandler(router) // cors

	return corsHandler
}

func corsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем запросы с любого источника
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Разрешаем определенные методы запросов
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE ,OPTIONS")

		// Разрешаем определенные заголовки запросов
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Если это предварительный запрос OPTIONS, просто отправляем пустой ответ
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Пропускаем запрос к следующему обработчику
		next.ServeHTTP(w, r)
	})
}
