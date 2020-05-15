package signaling

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Server struct{}

var room = NewRoom()

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Error(errors.Wrap(err, "problem upgrading to websockets"))
		return
	}

	client := NewPeer(conn)

	logrus.Debugf("connection established by %s", client.id)

	for {
		message, err := client.GetNextMessage()
		if err != nil {
			logrus.Error(errors.Wrap(err, "problem getting message from client"))

			// TODO add error handling
			return
		}

		logrus.Debugf("got message %v", message)

		room.Dispatch(client, message)
	}
}
