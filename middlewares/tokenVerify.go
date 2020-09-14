package middlewares

import (
	"github.com/gin-gonic/gin"
	"kecamatan_app/auth"
	"kecamatan_app/utils"
	"net/http"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request)
		if err != nil {
			utils.ErrorMessage(c, http.StatusUnauthorized, "You need to be authorized to access this route")
			c.Abort()
			return
		}
		c.Next()
	}
}
