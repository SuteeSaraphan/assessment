package expanse

import (
	"database/sql"
	"log"
	"os"
)

var db *sql.DB

const createTableSQL = `
CREATE TABLE IF NOT EXISTS expenses (
	id SERIAL PRIMARY KEY,
	title TEXT,
	amount FLOAT,
	note TEXT,
	tags TEXT[]
);`

func InitDB(dbUrl string) error {
	dbStr := os.Getenv("DATABASE_URL")
	// Connect DB
	if dbUrl == "" {
		dbUrl = dbStr
	}
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return err
	}

	// Create Table
	_, err = db.Exec(createTableSQL)
	if err != nil {
		return err
	}
	return nil
}

func setDB() *sql.DB {
	dbStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbStr)
	if err != nil {
		log.Fatal("Connection Fail")
	}
	return db
}
