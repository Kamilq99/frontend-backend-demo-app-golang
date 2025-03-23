package main

import (
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Konfiguracja CORS
	router.Use(cors.Default())

	staticPath := filepath.Join("..", "frontend")
	router.Static("/static", staticPath)

	router.GET("/text", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"text": "Hello World",
		})
	})

	router.Run(":8080")
}
