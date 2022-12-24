package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const defaultPort = ":2565"

func main() {
	e := echo.New()
	e.GET("/", func(s echo.Context) error {
		return s.String(200, "Hello")
	})

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	e.Logger.Fatal(e.Start(port))

}
