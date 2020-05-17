package signaling

import (
	"errors"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var ErrPeerNotFound = errors.New("peer not found")

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . RoomMember

type RoomMember interface {
	ID() PeerID
	SendMessage(Message) error
	Timedout() bool
	Close()
}

type Room struct {
	members     map[PeerID]RoomMember
	membersLock *sync.RWMutex
}

func NewRoom() *Room {
	return &Room{
		members:     map[PeerID]RoomMember{},
		membersLock: &sync.RWMutex{},
	}
}

func (r *Room) StartReaper(interval time.Duration) {
	go func() {
		for {
			logrus.Debugf("running reaper")

			r.membersLock.Lock()

			for _, member := range r.members {
				if member.Timedout() {
					member.Close()

					delete(r.members, member.ID())

					for _, peer := range r.members {
						message := Message{
							Type:          MessageTypeLeave,
							SourceID:      member.ID(),
							DestinationID: peer.ID(),
							Payload: map[string]interface{}{
								"reason": "timeout",
							},
						}

						err := peer.SendMessage(message)
						if err != nil {
							logrus.Warnf("problem broadcasting message to peer %s", peer.ID())
						}
					}
				}
			}

			r.membersLock.Unlock()

			time.Sleep(interval)
		}
	}()
}

func (r *Room) Dispatch(source RoomMember, message Message) {
	logrus.Debugf("Message type: %s", message.Type)

	switch message.Type {
	case MessageTypeJoin:
		r.AddMember(source)

		err := r.Broadcast(message)
		if err != nil {
			logrus.Error(err)
		}
	case MessageTypeLeave:
		r.RemoveMember(source)

		err := r.Broadcast(message)
		if err != nil {
			logrus.Error(err)
		}
	case MessageTypeOffer:
		err := r.MessageMember(message)
		if err != nil {
			logrus.Error(err)
		}
	case MessageTypeAnswer:
		err := r.MessageMember(message)
		if err != nil {
			logrus.Error(err)
		}
	case MessageTypeICECandidate:
		err := r.MessageMember(message)
		if err != nil {
			logrus.Error(err)
		}
	default:
		logrus.Warnf(`unknown message type %s`, message.Type)
		return
	}
}

func (r *Room) GetMember(peerID PeerID) RoomMember {
	r.membersLock.RLock()
	defer r.membersLock.RUnlock()

	member, ok := r.members[peerID]
	if !ok {
		return nil
	}

	return member
}

func (r *Room) AddMember(member RoomMember) {
	if r.GetMember(member.ID()) != nil {
		logrus.Warnf("member %s already present", member.ID())
		return // members already present
	}

	r.membersLock.Lock()
	defer r.membersLock.Unlock()

	r.members[member.ID()] = member

	logrus.Debugf("added member %s", member.ID())
}

func (r *Room) RemoveMember(members RoomMember) {
	r.membersLock.Lock()
	defer r.membersLock.Unlock()

	delete(r.members, members.ID())

	logrus.Debugf("removed members %s", members.ID())
}

func (r *Room) MessageMember(message Message) error {
	members := r.GetMember(message.DestinationID)
	if members == nil {
		logrus.Warnf("cannot find members %s", message.DestinationID)
		return nil // Don't error, just skip
	}

	err := members.SendMessage(message)
	if err != nil {
		logrus.Warnf("problem setting message to members %s", message.DestinationID)
		return nil
	}

	logrus.Debugf("sent members %s messsage %s", message.DestinationID, message)

	return nil
}

func (r *Room) Broadcast(message Message) error {
	r.membersLock.RLock()
	defer r.membersLock.RUnlock()

	logrus.Debugf("broadcasting message: %s", message)

	for _, member := range r.members {
		// Don't send messages to source
		if member.ID() == message.SourceID {
			continue
		}

		err := member.SendMessage(message)
		if err != nil {
			logrus.Warnf("problem broadcasting message to members %s", member.ID())
		}
	}

	return nil
}
