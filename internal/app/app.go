package app

import (
	"food_delivery/internal/config"
	"food_delivery/internal/storage"
	"log"
)

type App struct {
	Cfg    *config.Config
	Logger *log.Logger
	Store  *storage.Storage
}

func New(cfg *config.Config, logger *log.Logger, store *storage.Storage) *App {
	return &App{
		Cfg:    cfg,
		Logger: logger,
		Store:  store,
	}
}
