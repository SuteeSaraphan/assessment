package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/SuteeSaraphan/assessment/expanse"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

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

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")

		if authToken != "November 10, 2009" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authentication token")
		}
		return next(c)
	}
}

func main() {
	dbStr := os.Getenv("DATABASE_URL")

	e := echo.New()
	expanse.InitDB(dbStr)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Bonus middleware check Autorization
	e.Use(AuthMiddleware)

	// Routes
	e.GET("/", hello)
	e.GET("/expenses", expanse.GetAllExpensesHandler)
	e.GET("/expenses/:id", expanse.GetIdExpensesHandler)
	e.POST("/expenses", expanse.CreateExpenseHandler)
	e.PUT("/expenses/:id", expanse.UpdateExpenseHandler)

	go func() {
		if err := e.Start(os.Getenv("PORT")); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatalf("shutting down the server: %v", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
