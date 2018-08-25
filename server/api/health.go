package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckHandler provides a health check endpoint used to determine if the
// server is running
func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}
