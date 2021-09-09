package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func corsMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func Start(port int) {
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
	if port < 999 {
		port = 6080
	}
	var portString = ":" + strconv.Itoa(port)
	// Start and run the server
	router.Run(portString)
	fmt.Println("Server Started on port :", strconv.Itoa(port))
}
