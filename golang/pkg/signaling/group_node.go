package signaling

import "github.com/sirupsen/logrus"

type GroupDetails struct {
	ID             NodeID `json:"id"`
	Name           string `json:"name"`
	Limit          int    `json:"limit"`
	DependentCount int    `json:"num_members"`
}

type GroupNode struct {
	Node
	Dependents
}

func NewGroupNode(id NodeID, parent *Room, limit int) GroupNode {
	return GroupNode{
		Node:       NewNode(id, parent),
		Dependents: NewDependents(limit),
	}
}

func (g *GroupNode) Receive(message Message) {
	switch message.Type {
	case MessageTypeLeave:
		dependent := g.GetDependent(message.SourceID)
		if dependent == nil {
			logrus.Warnf("dependent not found %s", message.SourceID)
			return
		}

		dependent.SetParent(nil)
		g.RemoveDependent(dependent)

		logrus.Debugf("removed member %s from group %s", dependent.ID(), g.ID())
	case MessageTypeOffer, MessageTypeAnswer, MessageTypeICECandidate:
		g.MessageDependent(message)
	default:
		logrus.Warnf(`unknown message type %s`, message.Type)
		return
	}
}

func (g *GroupNode) GetDetails() GroupDetails {
	return GroupDetails{
		ID:             g.ID(),
		Name:           string(g.ID()),
		Limit:          g.Dependents.GetLimit(),
		DependentCount: g.Dependents.GetDependentsCount(),
	}
}
