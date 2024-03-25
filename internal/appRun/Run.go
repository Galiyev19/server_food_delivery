package apprun

import (
	"food_delivery/internal/app"
	"food_delivery/internal/config"
	"food_delivery/internal/data"
	"food_delivery/internal/server"
	"food_delivery/internal/storage/sqlite"
	"log"
	"os"
	"time"
)

func Run() {
	cfg := config.New()
	if err := cfg.InitConfig("config.json", cfg); err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := app.New(cfg, logger)

	storage, err := sqlite.New(app.Cfg.StoreDriver, app.Cfg.StorePath, app.Cfg.MigrationPath)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	rate := data.Rating{
		Rate:  4.5,
		Count: 120,
	}
	products := data.Products{
		ID:          1,
		CreadAt:     time.Now(),
		Title:       "test",
		Description: "Test test test test test",
		Category:    "test",
		Price:       100.0,
		Image:       "test",
		Rating:      rate,
	}

	_, err = storage.SaveProduct(products)

	log.Fatal(server.RunSever(app.Routes(), app.Cfg.Address))
}
