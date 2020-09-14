package utils

import "github.com/gin-gonic/gin"

// ErrorMessage ...
func ErrorMessage(c *gin.Context, status int, msg string) *gin.Context {
	c.JSON(status, gin.H{
		"Code": "1111",
		"Message": msg,
	})
	return c
}
