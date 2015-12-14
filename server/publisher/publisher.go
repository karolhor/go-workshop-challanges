package publisher

import (
	"bytes"
	"encoding/json"
	"github.com/karolhor/go-workshops-challange/common"
	"github.com/karolhor/go-workshops-challange/common/config"
	"gopkg.in/redis.v3"
	"net/http"
)

type (
	// Publisher interface
	Publisher interface {
		Publish(*message.Message) error
	}

	// JsonApiPublisher part
	JsonApiPublisher struct {
		ClientURL string
	}

	// RedisPublisher part
	RedisPublisher struct {
		redisClient *redis.Client
		channel     string
	}
)

func (p *JsonApiPublisher) Publish(msg *message.Message) error {

	var msgBody bytes.Buffer
	encoder := json.NewEncoder(&msgBody)
	if err := encoder.Encode(msg); err != nil {
		return err
	}

	_, err := http.Post(p.ClientURL, "application/json", &msgBody)

	return err
}

func (p *RedisPublisher) Publish(msg *message.Message) error {
	return p.redisClient.Publish("channel", msg.ToJSON()).Err()
}

// NewRedisPublisher creates RedisPublisher with default options with given address
func NewRedisPublisher(redisConfig *config.RedisConfig) *RedisPublisher {
	redisPublisher := &RedisPublisher{channel: redisConfig.PubSubChannel}

	redisPublisher.redisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := redisPublisher.redisClient.Ping().Err(); err != nil {
		panic(err)
	}

	return redisPublisher
}
