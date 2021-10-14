package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sggts04/urlshortener-go/handlers"
)

func main() {
	router := gin.Default()

	// Serve Frontend
	router.GET("/", handlers.ServeFrontend)

	// Register Long URL: Generate Short URL or Save Custom URL
	router.POST("/", handlers.RegisterLongURL)

	// Redirection to Long URL
	router.GET("/:id", handlers.RedirectToLongURL)

	router.Run("localhost:8000")
}
