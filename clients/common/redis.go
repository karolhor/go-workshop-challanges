package common

import (
	"github.com/karolhor/go-workshops-challange/common/config"
	"gopkg.in/redis.v3"
	"github.com/karolhor/go-workshops-challange/common"
	"github.com/labstack/gommon/log"
)

type (
	// RedisSubscriber part
	RedisSubscriber struct {
		redisClient *redis.Client
		channel     string
	}
)

func (rs *RedisSubscriber) Subscribe(channel string, msgChan chan<- *message.Message) {
	pubsub, err := rs.redisClient.Subscribe(channel)

	if err != nil {
		panic(err)
	}
	defer pubsub.Close()

	for {
		redisMsg, err := pubsub.ReceiveMessage()

		if err != nil {
			log.Println("Error on receiving msg from redis: %v", err)
		}

		msg, err := message.NewMessageFromJSON(redisMsg.Payload)
		if err != nil {
			log.Println("Error on parsing msg: %v", err)
		}

		log.Println(msg.ToJSON())
		msgChan <- msg
	}
}

// RedisSubscriber creates RedisSubscriber with default options with given address
func NewRedisSubscriber(redisConfig *config.RedisConfig) *RedisSubscriber {
	redisSubscriber := &RedisSubscriber{channel: redisConfig.PubSubChannel}

	redisSubscriber.redisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := redisSubscriber.redisClient.Ping().Err(); err != nil {
		panic(err)
	}

	return redisSubscriber
}
