package services

import (
	"github.com/gin-gonic/gin"
	"github.com/sggts04/urlshortener-go/handlers"
)

func NewURLShorteningService(handler *handlers.URLShorteningHandler) *gin.Engine {
	router := gin.Default()

	// Serve Static Files
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*.html")
	// Serve Frontend
	router.GET("/", handler.ServeFrontend)
	// Register Long URL: Generate Short URL or Save Custom URL
	router.POST("/", handler.RegisterLongURL)
	// Redirection to Long URL
	router.GET("/:id", handler.RedirectToLongURL)

	return router
}
