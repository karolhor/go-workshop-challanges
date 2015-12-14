package main

import (
	"github.com/karolhor/go-workshops-challange/clients/common"
	"github.com/karolhor/go-workshops-challange/clients/common/config"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"github.com/karolhor/go-workshops-challange/common"
)

type Message struct {
	ID      bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	message.Message
}

func insertMessageToMongo(msgs <-chan *message.Message) {
	msgFromChannel := <-msgs
	msg := &Message{Message:*msgFromChannel}

	msg.ID = bson.NewObjectId()

	err := messagesCollection.Insert(msg)
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

	var msgChannel = make(chan *message.Message)

	go insertMessageToMongo(msgChannel)
	rs.Subscribe(mongoConfig.RedisConfig.PubSubChannel, msgChannel)
}
