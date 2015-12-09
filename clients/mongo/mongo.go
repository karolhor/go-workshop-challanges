package main

import (
	"encoding/json"
	"github.com/karolhor/go-workshops-challange/clients/common"
	"github.com/karolhor/go-workshops-challange/clients/common/config"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/redis.v3"
	"log"
)

type Message struct {
	ID      bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Content string        `json:"message"`
	Owner   string        `json:"owner"`
}

func (m *Message) ToJSON() string {
	jsonMsg, err := json.Marshal(m)

	if err != nil {
		log.Println("Could not encode msg as json: %v", err)
	}

	return string(jsonMsg)
}

func NewMessageFromJSON(data string) (msg *Message, err error) {
	msg = &Message{ID: bson.NewObjectId()}

	err = json.Unmarshal([]byte(data), msg)

	return
}

func insertMessageToMongo(redisMsg *redis.Message) {
	msg, err := NewMessageFromJSON(redisMsg.Payload)

	if err != nil {
		log.Println("Not valid message content from redis: %s", redisMsg.Payload)
	}

	err = messagesCollection.Insert(msg)
	if err != nil {
		log.Println("Insert message error: %v", err)
		return
	}

	msgFromDb := &Message{}
	err = messagesCollection.FindId(msg.ID).One(msgFromDb)

	if err != nil {
		log.Println("Can't find document from mongo: %v", err)
	}

	log.Println(msgFromDb.ToJSON())
}

var messagesCollection *mgo.Collection

func main() {
	configPath := kingpin.Flag("config", "Configuration path").Short('c').Required().String()

	kingpin.Parse()
	mongoConfig := config.NewMongoConfigFromJSONFile(configPath)

	sess, err := mgo.Dial(mongoConfig.MongoDBConfig.URL)
	if err != nil {
		log.Fatalln("Can't connect to mongo, error %v", err)
	}
	defer sess.Close()

	messagesCollection = sess.DB(mongoConfig.MongoDBConfig.DbName).C("messages")

	log.Println("Start listening for redis msg on channel: " + mongoConfig.RedisConfig.PubSubChannel)

	rs := common.NewRedisSubscriber(mongoConfig.RedisConfig)
	rs.Subscribe(mongoConfig.RedisConfig.PubSubChannel, insertMessageToMongo)
}
