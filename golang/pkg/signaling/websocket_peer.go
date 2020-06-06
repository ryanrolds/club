package signaling

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lucsky/cuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . websocket.Conn

type PeerID string

const (
	timeout      = time.Second * 30
	pingInterval = time.Second * 20
)

type WebsocketPeer struct {
	Node

	messages chan Message

	heartbeat     time.Time
	heartbeatLock sync.Mutex

	conn     *websocket.Conn
	connLock sync.Mutex

	wait sync.WaitGroup
}

func NewWebsocketPeer(conn *websocket.Conn) *WebsocketPeer {
	return &WebsocketPeer{
		Node: NewNode(NodeID(cuid.New()), nil),

		heartbeat:     time.Now(),
		heartbeatLock: sync.Mutex{},

		conn:     conn,
		connLock: sync.Mutex{},

		wait: sync.WaitGroup{},
	}
}

func (p *WebsocketPeer) Receive(message Message) {
	p.messages <- message
}

func (p *WebsocketPeer) PumpWrite() {
	ticker := time.NewTicker(pingInterval)

	defer func() {
		ticker.Stop()
		p.conn.Close()
	}()

	for {
		select {
		case message, ok := <-p.messages:
			p.conn.SetWriteDeadline(time.Now().Add(timeout))

			err := p.conn.WriteJSON(message)
			if err != nil {
				logrus.Error(err)
				break
			}

			if !ok {
				// The hub closed the channel.
				p.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

		case <-ticker.C:
			p.conn.SetWriteDeadline(time.Now().Add(timeout))
			if err := p.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (p *WebsocketPeer) PumpRead() {
	defer func() {
		p.parent.Receive(NewLeaveMessage(p.ID()))
		p.conn.Close()
	}()

	p.conn.SetReadDeadline(time.Now().Add(timeout))
	p.conn.SetPongHandler(func(string) error {
		p.conn.SetReadDeadline(time.Now().Add(pingInterval))
		return nil
	})

	p.wait.Add(1)

	for {
		message, err := p.getNextMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				logrus.Warn(errors.Wrap(err, "unexpected close"))
			} else {
				logrus.Warn(errors.Wrap(err, "problem getting message from client"))
			}

			break
		}

		logrus.Debugf("got message %v", message)

		p.parent.Receive(message)
	}

	p.wait.Done()
}

func (p *WebsocketPeer) Close() {
	p.conn.Close()
}

func (p *WebsocketPeer) WaitForDisconnect() {
	p.wait.Wait()
}

func (p *WebsocketPeer) getNextMessage() (Message, error) {
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
