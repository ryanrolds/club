package signalling

import (
	"encoding/json"
	"log"

	"gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

type Message struct {
	Type          string                 `json:"type" validate:"required"`
	SourceID      PeerID                 `json:"peerId" validate:"required"`
	DestinationID PeerID                 `json:"destId"`
	Payload       map[string]interface{} `json:"payload" validate:"required"`
}

func NewMessageFromBytes(sourceID PeerID, data []byte) (Message, error) {
	var message Message
	err := json.Unmarshal(data, &message)
	if err != nil {
		log.Println(err)
		return Message{}, err
	}

	message.SourceID = sourceID

	log.Println(message)

	err = validate.Struct(message)
	if err != nil {
		log.Println(err.Error())
		return Message{}, err
	}

	return message, nil
}

func (m *Message) ToJSON() ([]byte, error) {
	b, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
		return []byte{}, err
	}

	return b, nil
}
