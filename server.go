package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const defaultPort = ":2565"
const createTableSQL = `
CREATE TABLE IF NOT EXISTS expenses (
	id SERIAL PRIMARY KEY,
	title TEXT,
	amount FLOAT,
	note TEXT,
	tags TEXT[]
);`

func initDB() error {
	// Connect DB
	dbStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbStr)
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

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Welcome to the server KKGO:assessment")
}

func main() {
	err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	//use middleware to log requests
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			e.Logger.Printf("Request: %s %s %s", c.Request().Method, c.Request().URL, c.Request().RemoteAddr)
			return next(c)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	errServer := e.Start(port)
	if err != nil {
		e.Logger.Fatal(errServer)
	}

}
