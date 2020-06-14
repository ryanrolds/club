package signaling

import (
	"sync"

	"github.com/sirupsen/logrus"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . ReceiverNode

type ReceiverNode interface {
	ID() NodeID
	SetParent(ReceiverNode)
	GetParent() ReceiverNode
	Receive(Message)
}

type MemberDetails struct {
	ID   NodeID `json:"id"`
	Name string `json:"name"`
}

type Members struct {
	limit       int
	members     map[NodeID]ReceiverNode
	membersLock *sync.RWMutex
}

func NewMembers(limit int) Members {
	return Members{
		limit:       limit,
		members:     map[NodeID]ReceiverNode{},
		membersLock: &sync.RWMutex{},
	}
}

func (c *Members) GetMember(id NodeID) ReceiverNode {
	c.membersLock.RLock()
	defer c.membersLock.RUnlock()

	member, ok := c.members[id]
	if !ok {
		return nil
	}

	return member
}

func (c *Members) GetMembersDetails() []MemberDetails {
	var details []MemberDetails

	return details
}

func (c *Members) GetLimit() int {
	return c.limit
}

func (c *Members) GetMemberCount() int {
	c.membersLock.RLock()
	defer c.membersLock.RUnlock()

	return len(c.members)
}

func (c *Members) AddMember(member ReceiverNode) {
	logrus.Debugf("adding member %s", member.ID())

	if c.GetMember(member.ID()) != nil {
		logrus.Warnf("member %s already present", member.ID())
		return // members already present
	}

	c.membersLock.RLock()
	defer c.membersLock.RUnlock()

	c.members[member.ID()] = member
	c.Broadcast(NewJoinMessage(member.ID()))
}

func (c *Members) RemoveMember(member ReceiverNode) {
	logrus.Debugf("removing member %s", member.ID())

	c.membersLock.RLock()
	defer c.membersLock.RUnlock()

	delete(c.members, member.ID())

	c.Broadcast(NewLeaveMessage(member.ID()))
}

func (c *Members) MessageMember(message Message) {
	member := c.GetMember(message.DestinationID)

	if member == nil {
		logrus.Warnf("cannot find member %s", message.DestinationID)
		return
	}

	member.Receive(message)

	logrus.Debugf("sent member %s messsage %s", member.ID(), message)
}

func (c *Members) Broadcast(message Message) {
	c.membersLock.RLock()
	defer c.membersLock.RUnlock()

	logrus.Debugf("broadcasting message: %s", message)

	for id, member := range c.members {
		// Don't send messages to source
		if id == message.SourceID {
			continue
		}

		member.Receive(message)
	}
}
