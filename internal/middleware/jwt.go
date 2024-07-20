package middleware

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/internal/util"
	"net/http"
)

func JWTAuth(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := util.ValidateJWT(c, secretKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
			})
			return
		}
		c.Next()
	}
}
