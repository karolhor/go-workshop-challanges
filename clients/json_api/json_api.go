package main

import (
	"fmt"
	"github.com/karolhor/go-workshops-challange/clients/common/config"
	"github.com/karolhor/go-workshops-challange/common"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"gopkg.in/alecthomas/kingpin.v2"
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

	msg := &message.Message{}
	err := c.Bind(msg)

	if err != nil {
		return echo.NewHTTPError(400, "body content is not in JSON format")
	}

	fmt.Println(msg.ToJSON())

	return c.NoContent(200)
}

func main() {
	configPath := kingpin.Flag("config", "Configuration path").Short('c').Required().String()

	kingpin.Parse()
	clientConfig := config.NewJsonApiConfigFromJSONFile(configPath)

	e := echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Routes
	e.Post("/", publishMessage)

	println("Running on port :" + clientConfig.Port)

	// Start server
	e.Run(":" + clientConfig.Port)
}
