package main

import (
	"github.com/karolhor/go-workshops-challange/clients/common"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/redis.v3"
	"log"
	"github.com/karolhor/go-workshops-challange/clients/common/config"
	"os"
	"io"
)

var Logger *log.Logger

func logMessage(msg *redis.Message) {
	Logger.Println(msg.Payload)
}

func main() {
	configPath := kingpin.Flag("config", "Configuration path").Short('c').Required().String()

	kingpin.Parse()
	loggerConfig := config.NewLoggerConfigFromJSONFile(configPath)

	logFile, err := os.OpenFile(loggerConfig.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open warning log file")
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(logFile, os.Stderr)
	Logger = log.New(multiWriter, "", log.Ldate|log.Ltime)

	Logger.Println("Start listening for redis msg on channel: " + loggerConfig.RedisConfig.PubSubChannel)

	rs := common.NewRedisSubscriber(loggerConfig.RedisConfig)
	rs.Subscribe(loggerConfig.RedisConfig.PubSubChannel, logMessage)
}
