package data

import (
	"database/sql"
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type URL struct {
	_id     int
	urlid   string
	longurl string
}

func GetLongURL(id string) (string, error) {
	row := db.QueryRow("SELECT * FROM urls WHERE urlid = ?", id)

	var res URL
	if err := row.Scan(&res._id, &res.urlid, &res.longurl); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("shorturl not found")
		}
		return "", errors.New("couldn't parse database result")
	}

	return res.longurl, nil
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

	row := db.QueryRow("SELECT * FROM urls WHERE urlid = ?", id)

	var res URL
	err = row.Scan(&res._id, &res.urlid, &res.longurl)
	if err == nil {
		// ID collision, it has already been generated and stored.
		return StoreLongURL(longURL, customId)
	} else if err != sql.ErrNoRows {
		return "", errors.New("couldn't parse database result")
	} else {
		_, err := db.Exec("INSERT INTO urls (urlid, longurl) VALUES (?, ?)", id, longURL)
		if err != nil {
			return "", errors.New("couldn't insert into database")
		}
		return id, nil
	}
}

func StoreCustomID(longURL string, customId string) (string, error) {
	row := db.QueryRow("SELECT * FROM urls WHERE urlid = ?", customId)

	var res URL
	err := row.Scan(&res._id, &res.urlid, &res.longurl)
	if err == nil {
		// ID collision, customId has already been stored.
		return "", errors.New("custom id already exists")
	} else if err != sql.ErrNoRows {
		return "", errors.New("couldn't parse database result")
	} else {
		_, err := db.Exec("INSERT INTO urls (urlid, longurl) VALUES (?, ?)", customId, longURL)
		if err != nil {
			return "", errors.New("couldn't insert into database")
		}
		return customId, nil
	}
}
