package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RecoveryHandler sends the correct error response when a handler panics
func RecoveryHandler(c *gin.Context, err interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err,
	})
}
