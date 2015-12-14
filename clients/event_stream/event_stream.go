package main

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/karolhor/go-workshops-challange/clients/common/config"
	"github.com/labstack/gommon/log"
	"github.com/karolhor/go-workshops-challange/clients/common"
	"github.com/karolhor/go-workshops-challange/common"
	"net/http"
)

var msgChannel = make(chan *message.Message)


func apiHandler(c *echo.Context) error {
	response := c.Response()

	response.Header().Set(echo.ContentType, "text/event-stream")
	response.WriteHeader(http.StatusOK)

	for {
		msg := <- msgChannel

		response.Write([]byte("data: "))
		response.Write([]byte(msg.ToJSON()))
		response.Write([]byte("\n\n"))
		response.Flush()
	}

	return nil
}

func main() {
	configPath := kingpin.Flag("config", "Configuration path").Short('c').Required().String()
	kingpin.Parse()

	eventStreamConfig := config.NewEventStreamConfigFromJSONFile(configPath)

	rs := common.NewRedisSubscriber(eventStreamConfig.RedisConfig)
	go rs.Subscribe(eventStreamConfig.RedisConfig.PubSubChannel, msgChannel)

	e := echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.Index(eventStreamConfig.StaticPath +"/index.html")
	e.Static("/", eventStreamConfig.StaticPath)
	e.Get("/api", apiHandler)

	log.Println("Event stream is running on port: %s", eventStreamConfig.Port)
	e.Run(":"+eventStreamConfig.Port)
}
