package repository

import (
	"go-boilerplate/internal/model"
)

type example struct {
}

var exampleList = []model.Example{}

type IExampleRepository interface {
	GetAll() []model.Example
	Create(request model.Example) error
}

func NewExampleRepository() IExampleRepository {
	return &example{}
}

func (r *example) GetAll() []model.Example {
	return exampleList
}

func (r *example) Create(request model.Example) error {
	exampleList = append(exampleList, request)
	return nil
}
