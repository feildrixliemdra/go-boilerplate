package handler

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/internal/payload"
	"go-boilerplate/internal/util"
	"net/http"
)

type exampleHandler struct {
	Option
}

type IExampleHandler interface {
	GetAll(c *gin.Context)
	Create(c *gin.Context)
}

func NewExampleHandler(option Option) IExampleHandler {
	return &exampleHandler{
		option,
	}
}

func (h *exampleHandler) GetAll(c *gin.Context) {
	list := h.Service.ExampleService.GetAll()
	util.GeneralSuccessResponse(c, "success get example list", list)
}

func (h *exampleHandler) Create(c *gin.Context) {

	var (
		req = payload.CreateExampleRequest{}
		err = c.ShouldBindJSON(&req)
	)

	if err != nil {
		util.ErrBindResponse(c, err)
		return
	}

	err = h.Service.ExampleService.Create(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal error",
		})

		return
	}

	util.GeneralSuccessResponse(c, "success create example", nil)

	return
}
