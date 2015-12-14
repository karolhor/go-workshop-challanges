package message

import (
	"encoding/json"
	"log"
)

// Message is basic struct shared with all clients
type Message struct {
	Msg   string `json:"message"`
	Owner string `json:"owner,omitempty"`
}

func (m *Message) ToJSON() string {
	jsonMsg, err := json.Marshal(m)

	if err != nil {
		log.Println("Could not encode msg as json: %v", err)
	}

	return string(jsonMsg)
}

func NewMessageFromJSON(data string) (msg *Message, err error) {
	msg = &Message{}

	err = json.Unmarshal([]byte(data), msg)

	return
}