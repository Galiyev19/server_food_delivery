package app

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *App) healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: available\n")
	fmt.Fprintf(w, "Enviroment: %s\n", app.Cfg.Enviroment)
	fmt.Fprintf(w, "Version: %s\n", app.Cfg.Version)
}

func (app *App) productListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ProductList\n")
}

func (app *App) showCategoryByName(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	categoryName := params.ByName("categoryName")

	fmt.Fprintf(w, "Category name: %s\n", categoryName)
}
