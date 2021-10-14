package data

import (
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

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
