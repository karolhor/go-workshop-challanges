package main

import (
	"github.com/karolhor/go-workshops-challange/clients/common"
	"github.com/karolhor/go-workshops-challange/clients/common/config"
	"gopkg.in/alecthomas/kingpin.v2"
	"io"
	"log"
	"os"
	"github.com/karolhor/go-workshops-challange/common"
)

var Logger *log.Logger

func logMessage(msgs <-chan *message.Message) {
	msgToLog := <- msgs
	Logger.Println(msgToLog.ToJSON())
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

	var msgChannel = make(chan *message.Message)

	go logMessage(msgChannel)
	rs.Subscribe(loggerConfig.RedisConfig.PubSubChannel, msgChannel)
}
