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

	MessageReceiver
}

type Dependents struct {
	limit          int
	dependents     map[NodeID]ReceiverNode
	dependentsLock *sync.RWMutex
}

func NewDependents(limit int) Dependents {
	return Dependents{
		limit:          limit,
		dependents:     map[NodeID]ReceiverNode{},
		dependentsLock: &sync.RWMutex{},
	}
}

func (c *Dependents) GetDependent(id NodeID) ReceiverNode {
	c.dependentsLock.RLock()
	defer c.dependentsLock.RUnlock()

	dependent, ok := c.dependents[id]
	if !ok {
		return nil
	}

	return dependent
}

func (c *Dependents) GetDependentsCount() int {
	c.dependentsLock.RLock()
	defer c.dependentsLock.RUnlock()

	return len(c.dependents)
}

func (c *Dependents) AddDependent(dependent ReceiverNode) {
	if c.GetDependent(dependent.ID()) != nil {
		logrus.Warnf("member %s already present", dependent.ID())
		return // members already present
	}

	c.dependentsLock.RLock()
	defer c.dependentsLock.RUnlock()

	c.dependents[dependent.ID()] = dependent

	logrus.Debugf("added member %s", dependent.ID())
}

func (c *Dependents) RemoveDependent(dependent ReceiverNode) {
	c.dependentsLock.RLock()
	defer c.dependentsLock.RUnlock()

	delete(c.dependents, dependent.ID())

	logrus.Debugf("removed members %s", dependent.ID())
}

func (c *Dependents) MessageDependent(message Message) {
	dependent := c.GetDependent(message.DestinationID)
	if dependent == nil {
		logrus.Warnf("cannot find dependent %s", message.DestinationID)
		return
	}

	dependent.Receive(message)

	logrus.Debugf("sent dependent %s messsage %s", message.DestinationID, message)
}

func (c *Dependents) Broadcast(message Message) {
	c.dependentsLock.RLock()
	defer c.dependentsLock.RUnlock()

	logrus.Debugf("broadcasting message: %s", message)

	for _, dependent := range c.dependents {
		// Don't send messages to source
		if dependent.ID() == message.SourceID {
			continue
		}

		dependent.Receive(message)
		logrus.Warnf("problem broadcasting message to members %s", dependent.ID())
	}
}
