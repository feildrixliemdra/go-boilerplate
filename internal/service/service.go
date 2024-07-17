package service

import (
	"go-boilerplate/internal/appcontext"
	"go-boilerplate/internal/repository"
)

type Service struct {
	ExampleService IExampleService
}

type Option struct {
	appcontext.Option
	Repository *repository.Repository
}

func InitiateService(option Option) *Service {

	exampleService := NewExampleService(option)

	return &Service{
		ExampleService: exampleService,
	}
}
