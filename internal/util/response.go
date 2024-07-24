package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-boilerplate/internal/constant"
	"go-boilerplate/internal/payload"
	val "go-boilerplate/internal/validator"
	"net/http"
)

func GeneralSuccessResponse(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, payload.Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrInternalResponse(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError,
		payload.Response{
			Success: false,
			Message: constant.InternalMessageErrorResponse,
			Error:   err,
		},
	)
}

func ErrBindResponse(c *gin.Context, err error) {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
			payload.Response{
				Success: false,
				Message: constant.ValidationFailureMessageResponse,
				Errors:  val.TranslateErrorValidator(err),
			})
		return
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, payload.Response{
		Success: false,
		Message: constant.BadRequestMessageResponse,
	})

	return
}

func ErrBadRequestResponse(c *gin.Context, err string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, payload.Response{
		Success: false,
		Message: constant.BadRequestMessageResponse,
		Error:   err,
	})

	return
}
