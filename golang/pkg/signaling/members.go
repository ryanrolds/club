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

func (m *Members) GetMember(id NodeID) ReceiverNode {
	m.membersLock.RLock()
	defer m.membersLock.RUnlock()

	member, ok := m.members[id]
	if !ok {
		return nil
	}

	return member
}

func (m *Members) GetMembersDetails() []MemberDetails {
	var details []MemberDetails

	for _, member := range m.members {
		details = append(details, MemberDetails{
			ID:   member.ID(),
			Name: string(member.ID()),
		})
	}

	return details
}

func (m *Members) GetLimit() int {
	return m.limit
}

func (m *Members) GetMembersCount() int {
	m.membersLock.RLock()
	defer m.membersLock.RUnlock()

	return len(m.members)
}

func (m *Members) AddMember(member ReceiverNode) {
	logrus.Debugf("adding member %s", member.ID())

	if m.GetMember(member.ID()) != nil {
		logrus.Warnf("member %s already present", member.ID())
		return // members already present
	}

	m.membersLock.RLock()
	defer m.membersLock.RUnlock()

	m.members[member.ID()] = member
	m.Broadcast(NewJoinMessage(member.ID()))
}

func (m *Members) RemoveMember(member ReceiverNode) {
	logrus.Debugf("removing member %s", member.ID())

	m.membersLock.RLock()
	defer m.membersLock.RUnlock()

	delete(m.members, member.ID())

	m.Broadcast(NewLeaveMessage(member.ID()))
}

func (m *Members) MessageMember(message Message) {
	member := m.GetMember(message.DestinationID)

	if member == nil {
		logrus.Warnf("cannot find member %s", message.DestinationID)
		return
	}

	member.Receive(message)

	logrus.Debugf("sent member %s messsage %s", member.ID(), message)
}

func (m *Members) Broadcast(message Message) {
	m.membersLock.RLock()
	defer m.membersLock.RUnlock()

	logrus.Debugf("broadcasting message: %s", message)

	for id, member := range m.members {
		// Don't send messages to source
		if id == message.SourceID {
			continue
		}

		member.Receive(message)
	}
}
