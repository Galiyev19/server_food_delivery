package app

import (
	"fmt"
	"food_delivery/internal/data"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *App) healthCheck(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"enviroment": app.Cfg.Enviroment,
			"version":    app.Cfg.Version,
		},
	}

	err := app.writeJson(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *App) productListHandler(w http.ResponseWriter, r *http.Request) {
	err := app.writeJson(w, http.StatusOK, nil, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *App) getProductByid(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	rate := data.Rating{
		Rate:  4.5,
		Count: 120,
	}
	products := data.Products{
		ID:       id,
		CreadAt:  time.Now(),
		Title:    "test",
		Category: "test",
		Price:    100.0,
		Image:    "test",
		Rating:   rate,
		Version:  1,
	}

	err = app.writeJson(w, http.StatusOK, products, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *App) createProduct(w http.ResponseWriter, r *http.Request) {
	type rate struct {
		Rate  float64 `json:"rate"`
		Count int64   `json:"count"`
	}

	type product struct {
		Title    string `json:"title"`
		Category string `json:"category"`
		Price    int64  `json:"price"`
		Image    string `json:"image"`
		Rating   rate   `json:"rating"`
	}

	var prod product

	err := app.readJson(w, r, &prod)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Fprintf(w, "%+v\n", prod)
}

func (app *App) showCategoryByName(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	categoryName := params.ByName("categoryName")

	fmt.Fprintf(w, "Category name: %s\n", categoryName)
}
