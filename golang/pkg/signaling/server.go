package signaling

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func NewUpgrader() websocket.Upgrader {
	return websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			// TODO: Validate env var list of origins here
			return true
		},
	}
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Dispatcher

type Dispatcher interface {
	Dispatch(RoomMember, Message) error
}

type Server struct {
	room     Dispatcher
	upgrader websocket.Upgrader
}

func NewServer(room Dispatcher) *Server {
	return &Server{
		room:     room,
		upgrader: NewUpgrader(),
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Error(errors.Wrap(err, "problem upgrading to websockets"))
		return
	}

	client := NewPeer(conn)

	logrus.Debugf("connection established by %s", client.ID())

	for {
		message, err := client.GetNextMessage()
		if err != nil {
			logrus.Warn(errors.Wrap(err, "problem getting message from client"))

			// TODO add error handling
			return
		}

		logrus.Debugf("got message %v", message)

		if message.Type == MessageTypeHeartbeat {
			client.Heartbeat()
			continue
		}

		err = s.room.Dispatch(client, message)
		if err != nil {
			orignal, jsonError := message.ToJSON()
			if jsonError != nil {
				logrus.WithError(err).Warnf("problem marshaling original message to JSON")
				continue
			}

			err = client.SendMessage(Message{
				Type:          MessageTypeError,
				SourceID:      PeerID("server"),
				DestinationID: client.ID(),
				Payload: MessagePayload{
					MessagePayloadKeyError:   err.Error(),
					MessagePayloadKeyMessage: string(orignal),
				},
			})
			if err != nil {
				logrus.WithError(err).Warnf("problem sending error message to client %s", client.ID())
				continue
			}
		}
	}
}
