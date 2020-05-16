package signaling

import (
	"errors"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const reaperInterval = time.Second * 15

var ErrPeerNotFound = errors.New("peer not found")

type Room struct {
	peers     map[PeerID]*Peer
	peersLock *sync.RWMutex
}

func NewRoom() *Room {
	return &Room{
		peers:     map[PeerID]*Peer{},
		peersLock: &sync.RWMutex{},
	}
}

func (r *Room) StartReaper() {
	go func() {
		for {
			logrus.Debugf("running reaper")

			r.peersLock.Lock()

			for _, member := range r.peers {
				if member.Timedout() {
					member.Close()

					delete(r.peers, member.ID)

					for _, peer := range r.peers {
						message := Message{
							Type:          MessageTypeLeave,
							SourceID:      member.ID,
							DestinationID: peer.ID,
							Payload: map[string]interface{}{
								"reason": "timeout",
							},
						}

						err := peer.SendMessage(message)
						if err != nil {
							logrus.Warnf("problem broadcasting message to peer %s", peer.ID)
						}
					}
				}
			}

			r.peersLock.Unlock()

			time.Sleep(reaperInterval)
		}
	}()
}

func (r *Room) Dispatch(source *Peer, message Message) {
	logrus.Debugf("Message type: %s", message.Type)

	switch message.Type {
	case MessageTypeJoin:
		r.AddPeer(source)

		err := r.Broadcast(message)
		if err != nil {
			logrus.Error(err)
		}
	case MessageTypeLeave:
		r.RemovePeer(source)

		err := r.Broadcast(message)
		if err != nil {
			logrus.Error(err)
		}
	case MessageTypeOffer:
		err := r.MessagePeer(message)
		if err != nil {
			logrus.Error(err)
		}
	case MessageTypeAnswer:
		err := r.MessagePeer(message)
		if err != nil {
			logrus.Error(err)
		}
	case MessageTypeICECandidate:
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
	r.peersLock.RLock()
	defer r.peersLock.RUnlock()

	peer, ok := r.peers[peerID]
	if !ok {
		return nil
	}

	return peer
}

func (r *Room) AddPeer(peer *Peer) {
	if r.GetPeer(peer.ID) != nil {
		logrus.Warnf("peer %s already present", peer.ID)
		return // Peer already present
	}

	r.peersLock.Lock()
	defer r.peersLock.Unlock()

	r.peers[peer.ID] = peer

	logrus.Debugf("added peer %s", peer.ID)
}

func (r *Room) RemovePeer(peer *Peer) {
	r.peersLock.Lock()
	defer r.peersLock.Unlock()

	delete(r.peers, peer.ID)

	logrus.Debugf("removed peer %s", peer.ID)
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
	r.peersLock.RLock()
	defer r.peersLock.RUnlock()

	logrus.Debugf("broadcasting message: %s", message)

	for _, peer := range r.peers {
		// Don't send messages to source
		if peer.ID == message.SourceID {
			continue
		}

		err := peer.SendMessage(message)
		if err != nil {
			logrus.Warnf("problem broadcasting message to peer %s", peer.ID)
		}
	}

	return nil
}
