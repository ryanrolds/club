package signaling

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

type MessageType string
type MessagePayloadKey string
type MessagePayload map[MessagePayloadKey]interface{}

const (
	MessageTypeHeartbeat    = MessageType("heartbeat")
	MessageTypeJoin         = MessageType("join")
	MessageTypeLeave        = MessageType("leave")
	MessageTypeOffer        = MessageType("offer")
	MessageTypeAnswer       = MessageType("answer")
	MessageTypeICECandidate = MessageType("icecandidate")
	MessageTypeError        = MessageType("error")
	MessageTypeKick         = MessageType("kick")
	MessageTypeShutdown     = MessageType("shutdown")

	MessagePayloadKeyGroup   = MessagePayloadKey("group")
	MessagePayloadKeyReason  = MessagePayloadKey("reason")
	MessagePayloadKeyError   = MessagePayloadKey("error")
	MessagePayloadKeyMessage = MessagePayloadKey("message")
)

var validate = validator.New()

type Message struct {
	Type          MessageType    `json:"type" validate:"required"`
	SourceID      NodeID         `json:"peerId" validate:"required"`
	DestinationID NodeID         `json:"destId"`
	Payload       MessagePayload `json:"payload" validate:"required"`
}

func NewJoinMessage(id NodeID) Message {
	return Message{
		Type:     MessageTypeJoin,
		SourceID: id,
		Payload:  MessagePayload{},
	}
}

func NewLeaveMessage(id NodeID) Message {
	return Message{
		Type:     MessageTypeLeave,
		SourceID: id,
		Payload:  MessagePayload{},
	}
}

func NewMessageFromBytes(sourceID NodeID, data []byte) (Message, error) {
	logrus.Debugf("Parsing message: %s", data)

	var message Message
	err := json.Unmarshal(data, &message)
	if err != nil {
		logrus.Error(err)
		return Message{}, err
	}

	message.SourceID = sourceID

	err = validate.Struct(message)
	if err != nil {
		logrus.Error(err)
		return Message{}, err
	}

	return message, nil
}

func (m *Message) ToJSON() ([]byte, error) {
	b, err := json.Marshal(m)
	if err != nil {
		logrus.Error(err)
		return []byte{}, err
	}

	return b, nil
}

func GetGroupIDFromMessage(message Message, defaultGroup NodeID) NodeID {
	groupString, ok := message.Payload[MessagePayloadKeyGroup]
	if !ok {
		return defaultGroup
	}

	return NodeID(groupString.(string))
}
