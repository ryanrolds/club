package signaling

import "github.com/sirupsen/logrus"

type GroupDetails struct {
	ID          NodeID          `json:"id"`
	Name        string          `json:"name"`
	Limit       int             `json:"limit"`
	MemberCount int             `json:"num_members"`
	Members     []MemberDetails `json:"members"`
}

type GroupNode struct {
	Node
	Members
}

func NewGroupNode(id NodeID, parent *Room, limit int) GroupNode {
	return GroupNode{
		Node:    NewNode(id, parent),
		Members: NewMembers(limit),
	}
}

func (g *GroupNode) Receive(message Message) {
	switch message.Type {
	case MessageTypeLeave:
		member := g.GetMember(message.SourceID)
		if member == nil {
			logrus.Warnf("member not found %s", message.SourceID)
			return
		}

		g.RemoveMember(member)
		member.SetParent(nil)

		logrus.Debugf("removed member %s from group %s", member.ID(), g.ID())
	case MessageTypeOffer, MessageTypeAnswer, MessageTypeICECandidate:
		g.MessageMember(message)
	default:
		logrus.Warnf(`unknown message type %s`, message.Type)
		return
	}
}

func (g *GroupNode) GetDetails() GroupDetails {
	return GroupDetails{
		ID:          g.ID(),
		Name:        string(g.ID()),
		Limit:       g.Members.GetLimit(),
		MemberCount: g.Members.GetMembersCount(),
		Members:     g.Members.GetMembersDetails(),
	}
}

func (g *GroupNode) AddMember(member ReceiverNode) {
	g.Members.AddMember(member)
	member.Receive(NewJoinedGroupMessage(member.ID(), g))
}

func (g *GroupNode) RemoveMember(member ReceiverNode) {
	g.Members.RemoveMember(member)
	member.Receive(NewLeftGroupMessage(member.ID(), g))
}
