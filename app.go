// main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Creates a router with default middleware:
	// logger and recovery (crash-free) middleware
	r := gin.Default()

	// Simple hello world route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World from Gin!",
			"status":  "success",
			"time":    c.Time,
		})
	})

	// Health check (good practice)
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Run on port 8080 (standard for most platforms)
	r.Run(":8080")
}
