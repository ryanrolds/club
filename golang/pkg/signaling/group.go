package signaling

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type GroupID string

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . GroupMember

type GroupMember interface {
	ID() PeerID
	SendMessage(Message) error
	Timedout() bool
	Close()
}

type Group struct {
	id          GroupID
	memberLimit int

	members     map[PeerID]GroupMember
	membersLock *sync.RWMutex
}

func NewGroup(group GroupID, limit int) *Group {
	return &Group{
		id:          group,
		memberLimit: limit,
		members:     map[PeerID]GroupMember{},
		membersLock: &sync.RWMutex{},
	}
}

func (g *Group) PruneStaleMembers() {
	g.membersLock.Lock()
	defer g.membersLock.Unlock()

	for _, member := range g.members {
		if member.Timedout() {
			member.Close()

			delete(g.members, member.ID())

			for _, peer := range g.members {
				message := Message{
					Type:          MessageTypeLeave,
					SourceID:      member.ID(),
					DestinationID: peer.ID(),
					Payload: map[MessagePayloadKey]interface{}{
						MessagePayloadKeyReason: "timeout",
					},
				}

				err := peer.SendMessage(message)
				if err != nil {
					logrus.Warnf("problem broadcasting message to peer %s", peer.ID())
				}
			}
		}
	}
}

func (g *Group) ID() GroupID {
	return g.id
}

func (g *Group) GetMember(peerID PeerID) GroupMember {
	g.membersLock.RLock()
	defer g.membersLock.RUnlock()

	member, ok := g.members[peerID]
	if !ok {
		return nil
	}

	return member
}

func (g *Group) GetMemberCount() int {
	g.membersLock.RLock()
	defer g.membersLock.RUnlock()

	return len(g.members)
}

func (g *Group) AddMember(member GroupMember) {
	if g.GetMember(member.ID()) != nil {
		logrus.Warnf("member %s already present", member.ID())
		return // members already present
	}

	g.membersLock.Lock()
	defer g.membersLock.Unlock()

	g.members[member.ID()] = member

	logrus.Debugf("added member %s", member.ID())
}

func (g *Group) RemoveMember(members GroupMember) {
	g.membersLock.Lock()
	defer g.membersLock.Unlock()

	delete(g.members, members.ID())

	logrus.Debugf("removed members %s", members.ID())
}

func (g *Group) MessageMember(message Message) error {
	member := g.GetMember(message.DestinationID)
	if member == nil {
		logrus.Warnf("cannot find members %s", message.DestinationID)
		return nil // Don't error, just skip
	}

	err := member.SendMessage(message)
	if err != nil {
		logrus.Warnf("problem setting message to members %s", message.DestinationID)
		return nil
	}

	logrus.Debugf("sent members %s messsage %s", message.DestinationID, message)

	return nil
}

func (g *Group) Broadcast(message Message) error {
	g.membersLock.RLock()
	defer g.membersLock.RUnlock()

	logrus.Debugf("broadcasting message: %s", message)

	for _, member := range g.members {
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
