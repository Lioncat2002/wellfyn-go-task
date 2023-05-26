package middlewares

import (
	"net/http"

	"main/utils/token"

	"github.com/gin-gonic/gin"
)

// JwtAuth middleware
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
