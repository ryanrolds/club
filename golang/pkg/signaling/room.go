package signaling

import (
	"errors"
	"sync"

	"github.com/sirupsen/logrus"
)

var ErrPeerNotFound = errors.New("peer not found")
var ErrGroupNotFound = errors.New("group not found")
var ErrGroupAlreadyExists = errors.New("group already exists")
var ErrMemberLacksGroup = errors.New("member lacks group")
var ErrInvalidMessageType = errors.New("invalid message type")
var ErrNonNilGroupRequired = errors.New("non-nil group required")

const (
	RoomDefaultID = NodeID("default")
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . NodeGroup

type ReceiverGroup interface {
	ReceiverNode

	AddDependent(ReceiverNode)
	GetDependent(NodeID) ReceiverNode
	RemoveDependent(ReceiverNode)

	Broadcast(Message)
	MessageDependent(Message)
}

type Room struct {
	Node
	Dependents

	groups     map[NodeID]ReceiverGroup
	groupsLock *sync.RWMutex
}

func NewRoom() *Room {
	return &Room{
		Node: NewNode("room", nil),

		// We need to hold peers that have joined the room, but not a group
		Dependents: NewDependents(0),

		groups: map[NodeID]ReceiverGroup{},
		// If we avoid creating groups, except for during startup, this mutex won't be needed
		groupsLock: &sync.RWMutex{},
	}
}

func (r *Room) Receive(message Message) {
	logrus.Debugf("Message type: %s", message.Type)

	member := r.GetDependent(message.SourceID)
	if member == nil {
		logrus.Warnf("member %s not found in room", member.ID())
		return
	}

	switch message.Type {
	case MessageTypeJoin:
		var group ReceiverGroup

		recevier := member.GetParent()
		if group != nil {
			group = r.GetGroup(recevier.ID())
			group.RemoveDependent(member)
			member.SetParent(nil)
		}

		groupID := GetGroupIDFromMessage(message, RoomDefaultID)
		group = r.GetGroup(groupID)
		if group == nil {
			return
		}

		member.SetParent(group)
		group.AddDependent(member)

		group.Broadcast(message)
	case MessageTypeLeave:
		receiver := member.GetParent()
		if receiver == nil {
			return
		}

		group := r.GetGroup(receiver.ID())
		group.RemoveDependent(member)

		group.Receive(message)
	case MessageTypeOffer, MessageTypeAnswer, MessageTypeICECandidate:
		group := member.GetParent()
		if group == nil {
			return
		}

		group.Receive(message)
	default:
		logrus.Warnf(`unknown message type %s`, message.Type)
		return
	}

	return
}

func (r *Room) AddGroup(group ReceiverGroup) error {
	if group == nil {
		return ErrNonNilGroupRequired
	}

	r.groupsLock.Lock()
	defer r.groupsLock.Unlock()

	_, ok := r.groups[group.ID()]
	if ok {
		return ErrGroupAlreadyExists
	}

	r.groups[group.ID()] = group

	return nil
}

func (r *Room) GetGroup(id NodeID) ReceiverGroup {
	r.groupsLock.RLock()
	defer r.groupsLock.RUnlock()

	group, ok := r.groups[id]
	if !ok {
		return nil
	}

	return group
}
