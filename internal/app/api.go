package app

import (
	"fmt"
	"net/http"
)

func (app *App) healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: available\n")
	fmt.Fprintf(w, "Enviroment: %s\n", app.Cfg.Enviroment)
	fmt.Fprintf(w, "Version: %s\n", app.Cfg.Version)
}
