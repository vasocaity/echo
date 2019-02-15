package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	//Echo instance
	e := echo.New()
	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// route to handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Vara")
	})
	//start server
	e.Logger.Fatal(e.Start(":3232"))
}
