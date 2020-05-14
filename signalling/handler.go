package signalling

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/lucsky/cuid"
	"gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type SignallingServer struct{}

type Message struct {
	Type          string                 `json:"type" validate:"required"`
	SourceID      string                 `json:"peerId" validate:"required"`
	DestinationID string                 `json:"destId"`
	Payload       map[string]interface{} `json:"payload" validate:"required"`
}

type Peer struct {
	id   string
	lock sync.Mutex
	conn *websocket.Conn
}

type Room map[string]*Peer

var room = Room{}
var roomLock = sync.RWMutex{}

func (s *SignallingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	sender := &Peer{
		id:   cuid.New(),
		lock: sync.Mutex{},
		conn: conn,
	}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))

		var parsed Message
		err = json.Unmarshal(p, &parsed)
		if err != nil {
			log.Println(err)
			return
		}

		parsed.SourceID = sender.id

		log.Println(parsed)

		err = validate.Struct(parsed)
		if err != nil {
			log.Println(err.Error())
			return
		}

		b, err := json.Marshal(parsed)
		if err != nil {
			log.Println(err)
			return
		}

		switch parsed.Type {
		case "heartbeat":
			log.Printf("heartbeat from %s", sender.id)
		case "join":
			// Add sender to joined peers
			roomLock.Lock()
			room[sender.id] = sender
			roomLock.Unlock()

			sendToPeers(sender, room, messageType, b)
		case "leave":
			sendToPeers(sender, room, messageType, b)
		case "offer":
			sendToPeer(parsed.DestinationID, room, messageType, b)
		case "answer":
			sendToPeer(parsed.DestinationID, room, messageType, b)
		case "icecandidate":
			sendToPeer(parsed.DestinationID, room, messageType, b)
		default:
			log.Printf(`unknown message type %s`, parsed.Type)
			return
		}
	}
}

func sendToPeer(destID string, room Room, messageType int, message []byte) {
	peer, ok := room[destID]
	if !ok {
		log.Println("could not find peer")
		return
	}

	peer.lock.Lock()
	err := peer.conn.WriteMessage(messageType, message)
	peer.lock.Unlock()

	if err != nil {
		log.Println(err)
		return
	}
}

func sendToPeers(sender *Peer, room Room, messageType int, message []byte) {
	for _, peer := range room {
		// Don't send messages to send
		if peer == sender {
			continue
		}

		peer.lock.Lock()
		err := peer.conn.WriteMessage(messageType, message)
		peer.lock.Unlock()

		if err != nil {
			log.Println(err)
			return
		}
	}
}
