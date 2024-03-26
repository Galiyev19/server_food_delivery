package app

import (
	"fmt"
	"food_delivery/internal/validator"
	"food_delivery/models"
	"net/http"

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

	fmt.Fprintf(w, "ID %d", id)

	err = app.writeJson(w, http.StatusOK, nil, nil)
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
		Title       string `json:"title"`
		Description string `json:"description"`
		Category    string `json:"category"`
		Price       int64  `json:"price"`
		Image       string `json:"image"`
		Rating      rate   `json:"rating"`
	}

	var p product

	err := app.readJson(w, r, &p)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	v := validator.New()

	// Copy the values from the input struct to a new Products struct.
	products := &models.Products{
		Title:       p.Title,
		Description: p.Description,
		Category:    p.Category,
		Price:       p.Price,
		Image:       p.Image,
		Rating:      models.Rating(p.Rating),
	}

	// Call the ValidateMovie() function and return a response containing the errors if any of the checks fail.
	if models.ValidatorProduct(v, products); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", p)
}

func (app *App) showCategoryByName(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	categoryName := params.ByName("categoryName")

	fmt.Fprintf(w, "Category name: %s\n", categoryName)
}
