package signaling

type GroupNode struct {
	Node
	Dependents
}

func NewGroupNode(id NodeID, parent *Room, limit int) *GroupNode {
	return &GroupNode{
		Node:       NewNode(id, parent),
		Dependents: NewDependents(limit),
	}
}

func (g *GroupNode) Receive(message Message) {
	if message.DestinationID != "" {
		g.Broadcast(message)
		return
	}

	g.MessageDependent(message)
}
