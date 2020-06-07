package signaling

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lucsky/cuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . WebsocketConn

type WebsocketConn interface {
	SetWriteDeadline(t time.Time) error
	SetReadDeadline(t time.Time) error
	SetPongHandler(h func(appData string) error)
	ReadMessage() (int, []byte, error)
	WriteMessage(messageType int, data []byte) error
	WriteJSON(interface{}) error
	Close() error
}

type PeerID string

const (
	timeout      = time.Second * 30
	pingInterval = time.Second * 20
)

type WebsocketPeer struct {
	Node
	messages chan Message
	conn     WebsocketConn
	wait     sync.WaitGroup
}

func NewWebsocketPeer(conn WebsocketConn, parent ReceiverNode) *WebsocketPeer {
	return &WebsocketPeer{
		Node:     NewNode(NodeID(cuid.New()), parent),
		messages: make(chan Message),
		conn:     conn,
		wait:     sync.WaitGroup{},
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

	err := p.conn.SetWriteDeadline(time.Now().Add(timeout))
	if err != nil {
		return
	}

	for {
		select {
		case message, ok := <-p.messages:
			err := p.conn.SetWriteDeadline(time.Now().Add(timeout))
			if err != nil {
				logrus.Error(err)
				return
			}

			if !ok {
				logrus.Warn("channel closed")
				err = p.conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					logrus.Warn(err)
				}

				return
			}

			logrus.Debugf("got message: %s", message)

			err = p.conn.WriteJSON(message)
			if err != nil {
				logrus.Error(err)
				return
			}
		case <-ticker.C:
			logrus.Debugf("ping %s", p.ID())

			err := p.conn.SetWriteDeadline(time.Now().Add(timeout))
			if err != nil {
				return
			}

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

	err := p.conn.SetReadDeadline(time.Now().Add(timeout))
	if err != nil {
		return
	}

	p.conn.SetPongHandler(func(string) error {
		logrus.Debugf("pong %s", p.ID())
		err := p.conn.SetReadDeadline(time.Now().Add(timeout))
		if err != nil {
			return err
		}

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
	err := p.conn.WriteMessage(websocket.CloseMessage, []byte{})
	if err != nil {
		logrus.Error(err)
	}

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
