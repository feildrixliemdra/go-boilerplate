package handler

import (
	"go-boilerplate/internal/config"
	"go-boilerplate/internal/service"
)

type Handler struct {
	UserHandler IUserHandler
}

func InitiateHandler(cfg *config.Config, services *service.Service) *Handler {
	return &Handler{
		UserHandler: NewUserHandler(cfg, services.UserService),
	}
}
