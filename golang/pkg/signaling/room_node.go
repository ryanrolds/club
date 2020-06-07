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
	DefaultGroupID = NodeID("default")
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . ReceiverGroup

type ReceiverGroup interface {
	ReceiverNode

	AddDependent(ReceiverNode)
	GetDependent(NodeID) ReceiverNode
	RemoveDependent(ReceiverNode)

	Broadcast(Message)
	MessageDependent(Message)
}

type Room struct {
	GroupNode

	groups     map[NodeID]ReceiverGroup
	groupsLock *sync.RWMutex
}

func NewRoom() *Room {
	return &Room{
		GroupNode: NewGroupNode("room", nil, 0),

		groups: map[NodeID]ReceiverGroup{},
		// If we avoid creating groups, except for during startup, this mutex won't be needed
		groupsLock: &sync.RWMutex{},
	}
}

func (r *Room) Receive(message Message) {
	logrus.Debugf("Message type: %s", message.Type)

	dependent := r.GetDependent(message.SourceID)
	if dependent == nil {
		logrus.Warnf("dependent %s not found in room", dependent.ID())
		return
	}

	switch message.Type {
	case MessageTypeJoin:
		var group ReceiverGroup

		// When joining a group, make sure to remove them from their previous group
		recevier := dependent.GetParent()
		if group != nil {
			group = r.GetGroup(recevier.ID())
			group.RemoveDependent(dependent)
			dependent.SetParent(nil)
		}

		groupID := GetGroupIDFromMessage(message, DefaultGroupID)
		group = r.GetGroup(groupID)
		if group == nil {
			return
		}

		dependent.SetParent(group)
		group.AddDependent(dependent)
		r.RemoveDependent(dependent)

		logrus.Debugf("added dependent %s to group %s", dependent.ID(), group.ID())
	case MessageTypeLeave:
		receiver := dependent.GetParent()
		if receiver == nil {
			return
		}

		group := r.GetGroup(receiver.ID())
		group.RemoveDependent(dependent)

		logrus.Debugf("removed dependent %s from group %s", dependent.ID(), group.ID())
	default:
		logrus.Warnf(`unknown message type %s`, message.Type)
		return
	}
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
