package common

import (
	"github.com/karolhor/go-workshops-challange/common/config"
	"gopkg.in/redis.v3"
)

// RedisSubscriber part
type RedisSubscriber struct {
	redisClient *redis.Client
	channel     string
}

type RedisReceiveMsgHandler func(*redis.Message)

func (rs *RedisSubscriber) Subscribe(channel string, handler RedisReceiveMsgHandler) {
	pubsub, err := rs.redisClient.Subscribe(channel)

	if err != nil {
		panic(err)
	}
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			panic(err)
		}

		handler(msg)
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
