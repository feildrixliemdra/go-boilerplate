package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func ErrBindResponse(c *gin.Context, err error) {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
			payload.Response{
				Success: false,
				Message: "validation failure",
				Errors:  val.TranslateErrorValidator(err),
			})
		return
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": err.Error(),
	})
	return
}
