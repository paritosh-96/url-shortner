package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func corsMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func Start() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	router.Use(corsMiddleWare())

	// health check api
	router.GET("/healthCheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Running",
		})
	})

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/actualUrl", GetActualURL)
		api.GET("/shortenUrl", ShortenURL)
	}

	// Start and run the server
	router.Run(":6080")
	fmt.Println("Server Started on port 6080")
}

