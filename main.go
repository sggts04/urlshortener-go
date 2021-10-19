package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sggts04/urlshortener-go/data"
	"github.com/sggts04/urlshortener-go/handlers"
)

func NewURLShorteningService() *gin.Engine {
	router := gin.Default()

	// Serve Frontend
	router.GET("/", handlers.ServeFrontend)
	// Register Long URL: Generate Short URL or Save Custom URL
	router.POST("/", handlers.RegisterLongURL)
	// Redirection to Long URL
	router.GET("/:id", handlers.RedirectToLongURL)

	return router
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASS := os.Getenv("MYSQL_PASS")
	MYSQL_ADDR := os.Getenv("MYSQL_ADDR")
	MYSQL_DB := os.Getenv("MYSQL_DB")

	err = data.InitDatabaseConnection(MYSQL_USER, MYSQL_PASS, MYSQL_ADDR, MYSQL_DB)
	if err != nil {
		log.Fatal("Coudn't connect to database")
	}

	service := NewURLShorteningService()
	service.Run("localhost:" + PORT)
}
