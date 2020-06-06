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

	conn     *websocket.Conn
	connLock sync.Mutex

	wait sync.WaitGroup
}

func NewWebsocketPeer(conn *websocket.Conn, parent ReceiverNode) *WebsocketPeer {
	return &WebsocketPeer{
		Node: NewNode(NodeID(cuid.New()), parent),

		messages: make(chan Message),

		conn:     conn,
		connLock: sync.Mutex{},

		wait: sync.WaitGroup{},
	}
}

func (p *WebsocketPeer) Receive(message Message) {
	logrus.Debugf("queueing message for %s", p.ID())

	p.messages <- message

	logrus.Debugf("finshed queuing message for %s", p.ID())
}

func (p *WebsocketPeer) PumpWrite() {
	ticker := time.NewTicker(pingInterval)

	p.wait.Add(1)

	defer func() {
		ticker.Stop()
		p.conn.Close()
		p.wait.Done()
		logrus.Debugf("exiting %s write pump", p.ID())
	}()

	p.conn.SetWriteDeadline(time.Now().Add(timeout))

	for {
		select {
		case message, ok := <-p.messages:
			p.conn.SetWriteDeadline(time.Now().Add(timeout))

			if !ok {
				logrus.Warn("channel closed")
				p.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			logrus.Debugf("got message: %s", message)

			err := p.conn.WriteJSON(message)
			if err != nil {
				logrus.Error(err)
				return
			}
		case <-ticker.C:
			logrus.Debugf("ping %s", p.ID())

			p.conn.SetWriteDeadline(time.Now().Add(timeout))

			if err := p.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				logrus.Debug(err)
				return
			}
		}
	}
}

func (p *WebsocketPeer) PumpRead() {
	p.wait.Add(1)

	defer func() {
		if p.parent != nil {
			p.parent.Receive(NewLeaveMessage(p.ID()))
		}

		p.conn.Close()
		p.wait.Done()
		logrus.Debugf("exiting %s read pump", p.ID())
	}()

	p.conn.SetReadDeadline(time.Now().Add(timeout))

	p.conn.SetPongHandler(func(string) error {
		logrus.Debugf("pong %s", p.ID())
		p.conn.SetReadDeadline(time.Now().Add(timeout))
		return nil
	})

	for {
		message, err := p.getNextMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				logrus.Warn(errors.Wrap(err, "unexpected close"))
			} else {
				logrus.Warn(errors.Wrapf(err, "problem getting message from client %s", p.ID()))
			}

			break
		}

		logrus.Debugf("got message %v", message)

		// Do not forward heartbeats to parent
		if message.Type == MessageTypeHeartbeat {
			continue
		}

		p.parent.Receive(message)
	}
}

func (p *WebsocketPeer) Close() {
	p.conn.WriteMessage(websocket.CloseMessage, []byte{})
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
