package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"food_delivery/internal/config"
	"food_delivery/internal/handlers"
	"food_delivery/internal/repository"
	"food_delivery/internal/service"
	sqlite "food_delivery/internal/storage"
)

func Run() {
	cfg := config.New() // create cfg and init
	if err := cfg.InitConfig("config.json", cfg); err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)                            // logger
	db, err := sqlite.CreateSqlDB(cfg.StoreDriver, cfg.StorePath, cfg.MigrationPath) // init db
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	r := repository.NewRepository(db)        // repo
	s := service.NewService(r)               // service
	h := handlers.NewHandler(s, cfg, logger) // handlers

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      h.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("starting server at http://localhost%s/v1/healthcheck\n", srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("ERROR: %w", err)
	}
}
