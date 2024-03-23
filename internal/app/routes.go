package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *App) Routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheck)
	router.HandlerFunc(http.MethodGet, "/v1/products", app.productListHandler)
	router.HandlerFunc(http.MethodGet, "/v1/product/:categoryName", app.showCategoryByName)
	return router
}
