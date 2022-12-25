package expanse

import (
	"database/sql"
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

func initDB(dbUrl string) error {
	// Connect DB

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return err
	}
	defer db.Close()
	// Create Table

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return err
	}
	return nil
}
