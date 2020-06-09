package signaling

type NodeID string

type Node struct {
	id     NodeID
	parent ReceiverNode
}

func NewNode(id NodeID, parent ReceiverNode) Node {
	return Node{
		id:     id,
		parent: parent,
	}
}

func (n *Node) ID() NodeID {
	return n.id
}

func (n *Node) SetParent(parent ReceiverNode) {
	n.parent = parent
}

func (n *Node) GetParent() ReceiverNode {
	return n.parent
}
