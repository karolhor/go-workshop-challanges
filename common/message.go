package message

import (
	"bytes"
	"encoding/json"
)

// Message is basic struct shared with all clients
type Message struct {
	Msg   string `json:"message"`
	Owner string `json:"owner,omitempty"`
}

func (msg Message) String() string {

	var msgResult bytes.Buffer
	encoder := json.NewEncoder(&msgResult)

	encoder.Encode(msg)

	return msgResult.String()
}
