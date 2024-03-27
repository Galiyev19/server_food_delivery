package handlers

import (
	"food_delivery/internal/config"
	"food_delivery/internal/service"
	"log"
)

type Handler struct {
	service *service.Service
	cfg     *config.Config
	logger  *log.Logger
}

func NewHandler(s *service.Service, config *config.Config, log *log.Logger) *Handler {
	return &Handler{
		service: s,
		cfg:     config,
		logger:  log,
	}
}
