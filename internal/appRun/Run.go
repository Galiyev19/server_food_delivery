package apprun

import (
	"food_delivery/internal/app"
	"food_delivery/internal/config"
	"food_delivery/internal/server"
	"food_delivery/internal/storage"
	"log"
	"os"
)

func Run() {
	cfg := config.New()
	if err := cfg.InitConfig("config.json", cfg); err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	db, err := storage.New(cfg.StoreDriver, cfg.StorePath, cfg.MigrationPath)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	store := storage.NewStorage(db)
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := app.New(cfg, logger, store)
	
	log.Fatal(server.RunSever(app.Routes(), app.Cfg.Address))
}
