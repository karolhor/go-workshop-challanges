package publisher

import (
	"bytes"
	"encoding/json"
	"github.com/karolhor/go-workshops-challange/common"
	"gopkg.in/redis.v3"
	"net/http"
)

// Publisher interface
type Publisher interface {
	Publish(*message.Message) error
}

// JsonApiPublisher part
type JsonApiPublisher struct {
	ClientURL string
}

func (p *JsonApiPublisher) Publish(msg *message.Message) error {

	var msgBody bytes.Buffer
	encoder := json.NewEncoder(&msgBody)
	if err := encoder.Encode(msg); err != nil {
		return err
	}

	_, err := http.Post(p.ClientURL, "application/json", &msgBody)

	return err
}

// RedisPublisher part
type RedisPublisher struct {
	RedisClient *redis.Client
}

func (p *RedisPublisher) Publish(msg *message.Message) error {
	return p.RedisClient.Publish("channel", msg.String()).Err()
}

// NewRedisPublisher creates RedisPublisher with default options with given address
func NewRedisPublisher(address string) *RedisPublisher {
	redisPublisher := &RedisPublisher{}
	redisPublisher.RedisClient = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := redisPublisher.RedisClient.Ping().Err(); err != nil {
		panic(err)
	}

	return redisPublisher
}
