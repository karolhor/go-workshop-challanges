package main

import (
	"github.com/karolhor/go-workshops-challange/clients/common"
	"github.com/karolhor/go-workshops-challange/common/config"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/redis.v3"
	"log"
)

func logMessage(msg *redis.Message) {
	log.Println(msg.Payload)
}

func main() {
	configPath := kingpin.Flag("config", "Configuration path").Short('c').Required().String()

	kingpin.Parse()
	loggerConfig := config.NewLoggerConfigFromJSONFile(configPath)

	println("Start listening for redis msg on channel: " + loggerConfig.RedisConfig.PubSubChannel)

	rs := common.NewRedisSubscriber(loggerConfig.RedisConfig)
	rs.Subscribe(loggerConfig.RedisConfig.PubSubChannel, logMessage)
}
