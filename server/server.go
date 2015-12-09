package main

import (
	"net/http"
	"strings"

	"github.com/karolhor/go-workshops-challange/common"
	"github.com/karolhor/go-workshops-challange/common/config"
	"github.com/karolhor/go-workshops-challange/server/publisher"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	publishers []publisher.Publisher
	serverConfig *config.ServerConfig
)

func init() {
	serverConfigPath := kingpin.Flag("config", "Server configuration path").Short('c').Required().String()

	kingpin.Parse()
	serverConfig = config.NewServerConfigFromJSONFile(serverConfigPath)

	publishers = []publisher.Publisher{
		&publisher.JsonApiPublisher{ClientURL: serverConfig.Clients.JSONApiUrl},
		publisher.NewRedisPublisher(serverConfig.RedisConfig)}
}

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

	msg := &message.Message{Owner: "Karol"}
	err := c.Bind(msg)

	if err != nil {
		return echo.NewHTTPError(400, "body content is not in JSON format")
	}

	for _, publisher := range publishers {
		go publisher.Publish(msg)
	}

	return c.JSON(http.StatusOK, msg)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Routes
	e.Post("/", publishMessage)

	println("Running on port :" + serverConfig.Port)

	// Start server
	e.Run(":" + serverConfig.Port)
}
