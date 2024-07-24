package handler

import (
	"go-boilerplate/internal/appcontext"
	"go-boilerplate/internal/service"
)

type Handler struct {
	UserHandler IUserHandler
}

type Option struct {
	appcontext.Option
	Service *service.Service
}

func InitiateHandler(opt Option) *Handler {
	return &Handler{
		UserHandler: NewUserHandler(opt),
	}
}
