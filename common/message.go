package message

import "encoding/json"

// Message is basic struct shared with all clients
type Message struct {
	Msg     string `json:"message"`
	AppName string `json:"app_name,omitempty"`
}

func (msg Message) String() string {
	result, _ := json.Marshal(msg)

	return string(result)
}
