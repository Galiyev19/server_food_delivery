package app

import "net/http"

func (app *App) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthCheck)
	return mux
}
