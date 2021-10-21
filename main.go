package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sggts04/urlshortener-go/data"
	"github.com/sggts04/urlshortener-go/handlers"
	"github.com/sggts04/urlshortener-go/services"
)

func main() {
	// Load ENV Vars
	if err := godotenv.Load(); err != nil {
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

	// Initialize Repository, Handler and Service (Using Dependency Injection)
	repo := data.NewRepository(db)
	handler := handlers.NewURLShorteningHandler(repo)
	service := services.NewURLShorteningService(handler)

	service.Run("localhost:" + PORT)
}
