package apprun

import (
	"food_delivery/internal/app"
	"food_delivery/internal/config"
	"food_delivery/internal/server"
	"log"
	"os"
)

func Run() {
	cfg := config.New()
	if err := cfg.InitConfig("config.json", cfg); err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := app.New(cfg, logger)
	log.Fatal(server.RunSever(app.Routes(), app.Cfg.Address))
}
