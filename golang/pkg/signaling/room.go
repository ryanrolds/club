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
	GetGroup() *Group
	SetGroup(*Group)

	GroupMember
}

type Room struct {
	groups     map[GroupID]Group
	groupsLock *sync.RWMutex
}

func NewRoom() *Room {
	return &Room{
		groups:     map[GroupID]Group{},
		groupsLock: &sync.RWMutex{},
	}
}

// Clients can disconenct without a leave event, iterate groups and tell them to
// remove stale members
func (r *Room) StartReaper(interval time.Duration) {
	go func() {
		for {
			logrus.Debugf("running reaper")

			r.groupsLock.RLock()

			for _, group := range r.groups {
				group.PruneStaleMembers()
			}

			r.groupsLock.RUnlock()

			time.Sleep(interval)
		}
	}()
}

func (r *Room) Dispatch(member RoomMember, message Message) {
	logrus.Debugf("Message type: %s", message.Type)

	switch message.Type {
	case MessageTypeJoin:
		group := member.GetGroup()
		if group != nil {
			group.RemoveMember(member)
			member.SetGroup(nil)
		}

		group = r.GetGroup(GetGroupIDFromMessage(message))
		member.SetGroup(group)
		group.AddMember(member)

		err := r.Broadcast(message)
		if err != nil {
			logrus.Error(err)
		}
	case MessageTypeLeave:
		group.RemoveMember(member)

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

func (r *Room) GetGroup() *Group {

}
