package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) userIdentity(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Получаем значение заголовка Authorization
		token := r.Header.Get(authorizationHeader)

		// Проверяем наличие токена
		if token == "" {
			// Если токен отсутствует, возвращаем ошибку "Unauthorized"
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Здесь вы можете выполнить дополнительную проверку на валидность токена,
		// например, проверить его наличие в базе данных или наличие подписи

		// Вызываем основной обработчик
		next(w, r, ps)
	}
}
