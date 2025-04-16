package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func HandleServiceError(c *gin.Context, err error) {
	switch {
	case strings.Contains(err.Error(), "is not allowed"):
		c.JSON(400, gin.H{
			"error":  "INVALID_FIELD",
			"detail": err.Error(),
		})
	case strings.Contains(err.Error(), "is protected"):
		c.JSON(403, gin.H{
			"error":  "PROTECTED_FIELD",
			"detail": err.Error(),
		})
	default:
		fmt.Print(err)
		c.JSON(500, gin.H{"error": "INTERNAL_ERROR"})
	}
}
