package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sggts04/urlshortener-go/data"
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

func main() {
	// Load ENV Vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASS := os.Getenv("MYSQL_PASS")
	MYSQL_ADDR := os.Getenv("MYSQL_ADDR")
	MYSQL_DB := os.Getenv("MYSQL_DB")

	// Init Database Connection
	db, err := data.InitDatabaseConnection(MYSQL_USER, MYSQL_PASS, MYSQL_ADDR, MYSQL_DB)
	if err != nil {
		log.Fatal("Coudn't connect to database")
	}

	// Initialize Repository, Handler and Serve (Using Dependency Injection)
	repo := data.NewRepository(db)
	handler := handlers.NewURLShorteningHandler(repo)
	service := NewURLShorteningService(handler)

	service.Run("localhost:" + PORT)
}
