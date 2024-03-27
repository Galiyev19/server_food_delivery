package handlers

import (
	"fmt"
	"net/http"
)

const (
	messageServerError   = "the server encountered a problem and could not process your request"
	messageNotFoundError = "the requsted resource could be found"
)

func (h *Handler) logError(r *http.Request, err error) {
	h.logger.Print(err)
}

// Return 500 status code
func (app *Handler) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}

	err := app.writeJson(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// Sever error if our applicatoion in runtime unexpected problem
func (app *Handler) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	app.errorResponse(w, r, http.StatusNotFound, messageServerError)
}

// Not Found
func (app *Handler) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	app.errorResponse(w, r, http.StatusNotFound, messageNotFoundError)
}

// Method not allowed
func (app *Handler) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

// Bad Request
func (app *Handler) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err)
}

// Note that the errors parameter here has the type map[string]string, which is exactly
// the same as the errors map contained in our Validator type.
func (app *Handler) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}
