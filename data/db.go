package data

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDatabaseConnection(MYSQL_USER string, MYSQL_PASS string, MYSQL_ADDR string, MYSQL_DB string) error {
	cfg := mysql.Config{
		User:                 MYSQL_USER,
		Passwd:               MYSQL_PASS,
		Net:                  "tcp",
		Addr:                 MYSQL_ADDR,
		DBName:               MYSQL_DB,
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return pingErr
	}
	log.Println("Connected to MySQL!")
	return nil
}
