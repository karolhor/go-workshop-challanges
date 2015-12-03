package main

import (
	"github.com/labstack/echo"
	"net/http"
)

// Handler
func hello(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func main() {
	e := echo.New()

	// Routes
	e.Get("/", hello)

	println("Running on port :8080")

	// Start server
	e.Run(":8080")
}
