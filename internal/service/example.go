package service

import (
	"go-boilerplate/internal/model"
	"go-boilerplate/internal/payload"
)

type example struct {
	Option
}

type IExampleService interface {
	GetAll() []model.Example
	Create(request payload.CreateExampleRequest) error
}

func NewExampleService(opt Option) IExampleService {
	return &example{
		opt,
	}
}

func (e *example) GetAll() []model.Example {
	return e.Repository.ExampleRepository.GetAll()
}

func (e *example) Create(request payload.CreateExampleRequest) error {
	return e.Repository.ExampleRepository.Create(model.Example{
		Name: request.Name,
		Age:  request.Age,
	})
}
