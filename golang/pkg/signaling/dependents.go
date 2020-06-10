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
	logrus.Debugf("adding dependent %s", dependent.ID())

	if c.GetDependent(dependent.ID()) != nil {
		logrus.Warnf("dependent %s already present", dependent.ID())
		return // members already present
	}

	c.dependentsLock.RLock()
	defer c.dependentsLock.RUnlock()

	c.dependents[dependent.ID()] = dependent

	c.Broadcast(NewJoinMessage(dependent.ID()))
}

func (c *Dependents) RemoveDependent(dependent ReceiverNode) {
	logrus.Debugf("removing dependent %s", dependent.ID())

	c.dependentsLock.RLock()
	defer c.dependentsLock.RUnlock()

	delete(c.dependents, dependent.ID())

	c.Broadcast(NewLeaveMessage(dependent.ID()))
}

func (c *Dependents) MessageDependent(message Message) {
	dependent := c.GetDependent(message.DestinationID)

	logrus.Error(message)

	if dependent == nil {
		logrus.Warnf("cannot find dependent %s", message.DestinationID)
		return
	}

	dependent.Receive(message)

	logrus.Debugf("sent dependent %s messsage %s", dependent.ID(), message)
}

func (c *Dependents) Broadcast(message Message) {
	c.dependentsLock.RLock()
	defer c.dependentsLock.RUnlock()

	logrus.Debugf("broadcasting message: %s", message)

	for id, dependent := range c.dependents {
		// Don't send messages to source
		if id == message.SourceID {
			continue
		}

		dependent.Receive(message)
	}
}
