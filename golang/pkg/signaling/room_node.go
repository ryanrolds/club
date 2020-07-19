package signaling

import (
	"errors"
	"sort"
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
	GetDetails() GroupDetails

	AddMember(ReceiverNode)
	GetMember(NodeID) ReceiverNode
	RemoveMember(ReceiverNode)

	Broadcast(Message)
	MessageMember(Message)
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

	member := r.GetMember(message.SourceID)
	if member == nil {
		logrus.Warnf("member %s not found in room", message.SourceID)
		return
	}

	switch message.Type {
	case MessageTypeJoin:
		var group ReceiverGroup

		// When joining a group, make sure to remove them from their previous group
		recevier := member.GetParent()
		if group != nil {
			group = r.GetGroup(recevier.ID())
			group.RemoveMember(member)
			member.SetParent(nil)
		}

		groupID := GetGroupIDFromMessage(message, DefaultGroupID)
		group = r.GetGroup(groupID)
		if group == nil {
			return
		}

		member.SetParent(group)
		group.AddMember(member)
		r.RemoveMember(member)

		logrus.Debugf("added member %s to group %s", member.ID(), group.ID())
	case MessageTypeLeave:
		r.RemoveMember(member)
		member.SetParent(nil)

		logrus.Debugf("removed member %s from room", member.ID())
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

func (r *Room) GetDetailsForGroups() []GroupDetails {
	r.groupsLock.RLock()
	defer r.groupsLock.RUnlock()

	var groups = make([]GroupDetails, 0)
	for _, group := range r.groups {
		groups = append(groups, group.GetDetails())
	}

	sort.Slice(groups, func(i, j int) bool {
		return groups[i].Name < groups[j].Name
	})

	return groups
}

func (r *Room) AddMember(member ReceiverNode) {
	r.GroupNode.Members.AddMember(member)
	member.Receive(NewJoinedRoomMessage(member.ID(), r))
}

func (r *Room) RemoveMember(member ReceiverNode) {
	r.GroupNode.Members.RemoveMember(member)
	member.Receive(NewLeftRoomMessage(member.ID(), r))
}
