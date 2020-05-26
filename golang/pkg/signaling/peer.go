package signaling

import (
	"sync"
	"time"

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

const (
	timeout = time.Second * 30
)

type Peer struct {
	id    PeerID
	group *Group

	heartbeat     time.Time
	heartbeatLock sync.Mutex

	conn     PeerConnection
	connLock sync.Mutex
}

func NewPeer(conn PeerConnection) *Peer {
	return &Peer{
		id:    PeerID(cuid.New()),
		group: nil,

		heartbeat:     time.Now(),
		heartbeatLock: sync.Mutex{},

		conn:     conn,
		connLock: sync.Mutex{},
	}
}

func (p *Peer) ID() PeerID {
	return p.id
}

func (p *Peer) GetGroup() *Group {
	return p.group
}

func (p *Peer) SetGroup(group *Group) {
	p.group = group
}

func (p *Peer) Heartbeat() {
	p.heartbeatLock.Lock()
	defer p.heartbeatLock.Unlock()

	p.heartbeat = time.Now()
}

func (p *Peer) Timedout() bool {
	p.heartbeatLock.Lock()
	defer p.heartbeatLock.Unlock()

	return p.heartbeat.Before(time.Now().Add(-timeout))
}

func (p *Peer) GetNextMessage() (Message, error) {
	_, data, err := p.conn.ReadMessage()

	if err != nil {
		return Message{}, err
	}

	message, err := NewMessageFromBytes(p.ID(), data)
	if err != nil {
		return Message{}, err
	}

	return message, nil
}

func (p *Peer) SendMessage(message Message) error {
	p.connLock.Lock()
	defer p.connLock.Unlock()

	err := p.conn.WriteJSON(message)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (p *Peer) Close() {
	p.connLock.Lock()
	defer p.connLock.Unlock()

	err := p.conn.Close()
	if err != nil {
		logrus.Error(err)
	}
}
