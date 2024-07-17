package handler

import (
	"go-boilerplate/internal/appcontext"
	"go-boilerplate/internal/service"
)

type Handler struct {
	ExampleHandler IExampleHandler
}

type Option struct {
	appcontext.Option
	Service *service.Service
}

func InitiateHandler(opt Option) *Handler {

	example := NewExampleHandler(opt)

	return &Handler{
		ExampleHandler: example,
	}
}
