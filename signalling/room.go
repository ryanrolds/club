package signalling

import (
	"errors"
	"log"
	"sync"
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
	switch message.Type {
	case "heartbeat":
		log.Printf("heartbeat from %s", source.id)
	case "join":
		r.AddPeer(source)

		err := r.Broadcast(message)
		if err != nil {
			log.Println(err)
		}
	case "leave":
		r.RemovePeer(source)

		err := r.Broadcast(message)
		if err != nil {
			log.Println(err)
		}
	case "offer":
		err := r.MessagePeer(message)
		if err != nil {
			log.Println(err)
		}
	case "answer":
		err := r.MessagePeer(message)
		if err != nil {
			log.Println(err)
		}
	case "icecandidate":
		err := r.MessagePeer(message)
		if err != nil {
			log.Println(err)
		}
	default:
		log.Printf(`unknown message type %s`, message.Type)
		return
	}
}

func (r *Room) GetPeer(peerId PeerID) *Peer {
	peer, ok := r.peers[peerId]
	if !ok {
		return nil
	}

	return peer
}

func (r *Room) AddPeer(peer *Peer) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()

	if r.GetPeer(peer.id) == nil {
		return // Peer already present
	}

	r.peers[peer.id] = peer
}

func (r *Room) RemovePeer(peer *Peer) {
	r.rwLock.Lock()
	delete(r.peers, peer.id)
	r.rwLock.Unlock()
}

func (r *Room) MessagePeer(message Message) error {
	peer := r.GetPeer(message.DestinationID)
	if peer != nil {
		log.Printf("cannot find peer %s", message.DestinationID)
		return nil // Don't error, just skip
	}

	err := peer.SendMessage(message)
	if err == nil {
		log.Printf("problem setting message to peer %s", message.DestinationID)
	}

	return nil
}

func (r *Room) Broadcast(message Message) error {
	for _, peer := range room.peers {
		// Don't send messages to source
		if peer.id == message.SourceID {
			continue
		}

		err := peer.SendMessage(message)
		if err != nil {
			log.Printf("problem broadcasting message to peer %s", message.DestinationID)
		}
	}

	return nil
}
