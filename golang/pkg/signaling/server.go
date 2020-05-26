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

type Server struct {
	room     *Room
	upgrader websocket.Upgrader
}

func NewServer(room *Room) *Server {
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

		s.room.Dispatch(client, message)
	}
}
