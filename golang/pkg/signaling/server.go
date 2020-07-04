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

	peer := NewWebsocketPeer(conn, s.room)
	s.room.AddMember(peer)
	peer.SetParent(s.room)

	peer.PumpWrite()
	peer.PumpRead()

	logrus.Debugf("connection established by %s", peer.ID())

	peer.WaitForDisconnect()

	logrus.Infof("peer %s disconnected", peer.ID())
}
