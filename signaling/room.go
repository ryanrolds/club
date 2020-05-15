package signaling

import (
	"errors"
	"sync"

	"github.com/sirupsen/logrus"
)

var ErrPeerNotFound = errors.New("peer not found")

type Room struct {
	peers  map[PeerID]*Peer
	rwLock *sync.RWMutex
}

func NewRoom() *Room {
	return &Room{
		peers:  map[PeerID]*Peer{},
		rwLock: &sync.RWMutex{},
	}
}

func (r *Room) Dispatch(source *Peer, message Message) {
	logrus.Debugf("Message type: %s", message.Type)

	switch message.Type {
	case "heartbeat":
		logrus.Infof("heartbeat from %s", source.id)
	case "join":
		r.AddPeer(source)

		err := r.Broadcast(message)
		if err != nil {
			logrus.Error(err)
		}
	case "leave":
		r.RemovePeer(source)

		err := r.Broadcast(message)
		if err != nil {
			logrus.Error(err)
		}
	case "offer":
		err := r.MessagePeer(message)
		if err != nil {
			logrus.Error(err)
		}
	case "answer":
		err := r.MessagePeer(message)
		if err != nil {
			logrus.Error(err)
		}
	case "icecandidate":
		err := r.MessagePeer(message)
		if err != nil {
			logrus.Error(err)
		}
	default:
		logrus.Warnf(`unknown message type %s`, message.Type)
		return
	}
}

func (r *Room) GetPeer(peerID PeerID) *Peer {
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()

	peer, ok := r.peers[peerID]
	if !ok {
		return nil
	}

	return peer
}

func (r *Room) AddPeer(peer *Peer) {
	if r.GetPeer(peer.id) != nil {
		logrus.Warnf("peer %s already present", peer.id)
		return // Peer already present
	}

	r.rwLock.Lock()
	defer r.rwLock.Unlock()

	r.peers[peer.id] = peer

	logrus.Debugf("added peer %s", peer.id)
}

func (r *Room) RemovePeer(peer *Peer) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()

	delete(r.peers, peer.id)

	logrus.Debugf("removed peer %s", peer.id)
}

func (r *Room) MessagePeer(message Message) error {
	peer := r.GetPeer(message.DestinationID)
	if peer == nil {
		logrus.Warnf("cannot find peer %s", message.DestinationID)
		return nil // Don't error, just skip
	}

	err := peer.SendMessage(message)
	if err != nil {
		logrus.Warnf("problem setting message to peer %s", message.DestinationID)
		return nil
	}

	logrus.Debugf("sent peer %s messsage %s", message.DestinationID, message)

	return nil
}

func (r *Room) Broadcast(message Message) error {
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()

	logrus.Debugf("broadcasting message: %s", message)

	for _, peer := range room.peers {
		// Don't send messages to source
		if peer.id == message.SourceID {
			continue
		}

		err := peer.SendMessage(message)
		if err != nil {
			logrus.Warnf("problem broadcasting message to peer %s", message.DestinationID)
		}
	}

	return nil
}
