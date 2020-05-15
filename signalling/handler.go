package signalling

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type SignallingServer struct{}

var room = NewRoom()

func (s *SignallingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := NewPeer(conn)

	for {
		message, err := client.GetNextMessage()
		if err != nil {
			log.Println(err)

			// TODO add error handling
			return
		}

		room.Dispatch(client, message)
	}
}
