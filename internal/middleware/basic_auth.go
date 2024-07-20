package middleware

import "github.com/gin-gonic/gin"

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"foo":   "bar",
		"test":  "1234",
		"admin": "1234",
	})
}
