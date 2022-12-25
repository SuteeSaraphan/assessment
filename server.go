package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/SuteeSaraphan/assessment/expanse"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
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

func InitDB(dbUrl string) error {
	// Connect DB

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

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Welcome to the server KKGO:assessment")
}

func main() {
	dbStr := os.Getenv("DATABASE_URL")

	e := echo.New()
	expanse.InitDB(dbStr)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/expenses", expanse.GetAllExpensesHandler)
	e.POST("/expenses", expanse.CreateExpenseHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := e.Start(port)
	if err != nil {
		e.Logger.Fatal(err)
	}

}
