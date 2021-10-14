package data

import (
	"context"
	"errors"
	"log"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client *mongo.Client
	ctx    context.Context
	cancel context.CancelFunc
}

var DB Database

func InitDatabaseConnection(connectionURI string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
		cancel()
		return err
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
		cancel()
		return err
	}

	log.Println("Connected to MongoDB!")
	DB = Database{client: client, ctx: ctx, cancel: cancel}
	return nil
}

func DisconnectDatabase() {
	DB.cancel()
	DB.client.Disconnect(DB.ctx)
	log.Println("Disconnected from MongoDB!")
}

var urls = map[string]string{}

func GetLongURL(id string) (string, error) {
	longURL, ok := urls[id]
	if !ok {
		// ID not found.
		return "", errors.New("short url not found")
	}
	return longURL, nil
}

func StoreLongURL(longURL string, customId string) (string, error) {
	if customId != "" {
		return StoreCustomID(longURL, customId)
	}
	id, err := gonanoid.New(6)
	if err != nil {
		// ID couldn't be generated.
		return "", errors.New("short url couldn't be generated")
	}
	if _, ok := urls[id]; ok {
		// ID collision, it has already been generated and stored.
		return StoreLongURL(longURL, customId)
	}
	urls[id] = longURL
	return id, nil
}

func StoreCustomID(longURL string, customId string) (string, error) {
	if _, ok := urls[customId]; ok {
		// ID collision, customId has already been stored.
		return "", errors.New("custom id already exists")
	}
	urls[customId] = longURL
	return customId, nil
}
