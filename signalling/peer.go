package signalling

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/lucsky/cuid"
)

type PeerID string

type Peer struct {
	id   PeerID
	lock sync.Mutex
	conn *websocket.Conn
}

func NewPeer(conn *websocket.Conn) *Peer {
	return &Peer{
		id:   PeerID(cuid.New()),
		lock: sync.Mutex{},
		conn: conn,
	}
}

func (p *Peer) GetNextMessage() (Message, error) {
	_, data, err := p.conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return Message{}, err
	}

	message, err := NewMessageFromBytes(p.id, data)
	if err != nil {
		return Message{}, err
	}

	log.Println(string(data))

	return message, nil
}

func (p *Peer) SendMessage(message Message) error {
	p.lock.Lock()
	err := p.conn.WriteJSON(message)
	p.lock.Unlock()

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
