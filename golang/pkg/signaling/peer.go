package signaling

import (
	"sync"

	"github.com/lucsky/cuid"
	"github.com/sirupsen/logrus"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . PeerConnection

type PeerConnection interface {
	ReadMessage() (int, []byte, error)
	WriteJSON(interface{}) error
	Close() error
}

type PeerID string

type Peer struct {
	ID   PeerID
	lock sync.Mutex
	conn PeerConnection
}

func NewPeer(conn PeerConnection) *Peer {
	return &Peer{
		ID:   PeerID(cuid.New()),
		lock: sync.Mutex{},
		conn: conn,
	}
}

func (p *Peer) GetNextMessage() (Message, error) {
	_, data, err := p.conn.ReadMessage()
	if err != nil {
		logrus.Error(err)
		return Message{}, err
	}

	message, err := NewMessageFromBytes(p.ID, data)
	if err != nil {
		return Message{}, err
	}

	return message, nil
}

func (p *Peer) SendMessage(message Message) error {
	p.lock.Lock()
	err := p.conn.WriteJSON(message)
	p.lock.Unlock()

	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *Peer) Leave() {
	p.lock.Lock()
	defer p.lock.Unlock()

	err := p.conn.Close()
	if err != nil {
		logrus.Error(err)
	}
}
