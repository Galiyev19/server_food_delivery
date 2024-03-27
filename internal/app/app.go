package app

import (
	"fmt"
	"food_delivery/internal/config"
	"food_delivery/internal/handlers"
	"food_delivery/internal/repository"
	"food_delivery/internal/service"
	sqlite "food_delivery/internal/storage"
	"log"
	"net/http"
	"os"
	"time"
)

func Run() {
	cfg := config.New()
	if err := cfg.InitConfig("config.json", cfg); err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	db, err := sqlite.CreateSqlDB(cfg.StoreDriver, cfg.StorePath, cfg.MigrationPath)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	// err = sqlite.InsertTestDataInUser(cfg.StoreDriver, cfg.StorePath, cfg.InitTest)
	// if err != nil {
	// 	log.Fatalf("ERROR: %v", err)
	// }

	r := repository.NewRepository(db)
	s := service.NewService(r)
	h := handlers.NewHandler(s, cfg, logger)

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
