package server

import (
	"fmt"
	"net/http"
	"time"
)

func RunSever(mux http.Handler, port string) error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("starting server at http://localhost%s\n", srv.Addr)
	err := srv.ListenAndServe()
	return err
}
