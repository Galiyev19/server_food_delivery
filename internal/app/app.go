package app

import (
	"food_delivery/internal/config"
	"log"
)

type App struct {
	Cfg *config.Config
	Log *log.Logger
}

func New(cfg *config.Config, logger *log.Logger) *App {
	return &App{
		Cfg: cfg,
		Log: logger,
	}
}
