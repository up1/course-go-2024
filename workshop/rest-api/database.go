package api

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const dbMaxIdleConns = 4
const dbMaxConns = 100

func NewDatabaseConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DB_URL"))

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(dbMaxConns)
	db.SetMaxIdleConns(dbMaxIdleConns)

	if err := db.Ping(); err != nil {
		log.Printf("Error pinging database: %v", err)
		return nil, err
	}
	log.Println("Database connection established")

	return db, nil
}
