package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func RunSever(router *httprouter.Router, port string) error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("starting server at http://localhost%s/v1/healthcheck\n", srv.Addr)
	err := srv.ListenAndServe()
	return err
}
