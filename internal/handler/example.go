package handler

import (
	"github.com/gin-gonic/gin"
)

type exampleHandler struct {
	Option
}

type IExampleHandler interface {
	GetAll(c *gin.Context)
}

func NewExampleHandler(option Option) IExampleHandler {
	return &exampleHandler{
		option,
	}
}

func (h *exampleHandler) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{"data": h.Service.ExampleService.GetAll()})
}
