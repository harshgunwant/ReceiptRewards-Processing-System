package utils

import (
	"github.com/gin-gonic/gin"
)

// SendError sends an error response in JSON format
func SendError(c *gin.Context, statusCode int, message string, details string) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"details": details,
	})
}
