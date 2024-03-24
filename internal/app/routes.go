package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *App) Routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowed)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheck)
	// router.HandlerFunc(http.MethodGet, "/v1/products", app.productListHandler)
	router.HandlerFunc(http.MethodPost, "/v1/products", app.createProduct)
	router.HandlerFunc(http.MethodGet, "/v1/products/:id", app.getProductByid)
	router.HandlerFunc(http.MethodGet, "/v1/product/:categoryName", app.showCategoryByName)
	return router
}
