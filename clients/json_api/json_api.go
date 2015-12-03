package main

import (
	"fmt"
	"github.com/karolhor/go-workshops-challange/common"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"net/http"
	"strings"
)

func assertContentTypeJSON(r *http.Request) *echo.HTTPError {
	ct := r.Header.Get(echo.ContentType)

	if !strings.HasPrefix(ct, echo.ApplicationJSON) {
		return echo.NewHTTPError(http.StatusBadRequest, "request: allowed Content-Type is 'application/json' only")
	}

	return nil
}

// Handler
func publishMessage(c *echo.Context) error {

	if err := assertContentTypeJSON(c.Request()); err != nil {
		return err
	}

	msg := &message.Message{AppName: "Karol"}
	err := c.Bind(msg)

	if err != nil {
		return echo.NewHTTPError(400, "body content is not in JSON format")
	}

	fmt.Println(msg.String())

	return c.NoContent(200)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Routes
	e.Post("/", publishMessage)

	println("Running on port :8000")

	// Start server
	e.Run(":8000")
}
