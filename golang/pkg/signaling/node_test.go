package signaling_test

import (
	"github.com/ryanrolds/club/pkg/signaling"
	"github.com/ryanrolds/club/pkg/signaling/signalingfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Node", func() {
	var (
		fakeParent *signalingfakes.FakeReceiverNode
		node       signaling.Node
	)

	BeforeEach(func() {
		fakeParent = &signalingfakes.FakeReceiverNode{}
		node = signaling.NewNode(signaling.NodeID("123"), fakeParent)
	})

	Context("NewNode", func() {
		It("should return node", func() {
			node = signaling.NewNode(signaling.NodeID("123"), fakeParent)
			Expect(node).To(BeAssignableToTypeOf(signaling.Node{}))
		})
	})

	Context("ID", func() {
		It("should return node id", func() {
			Expect(node.ID()).To(Equal(signaling.NodeID("123")))
		})
	})

	Context("GetParent", func() {
		It("should return node parent", func() {
			Expect(node.GetParent()).To(Equal(fakeParent))
		})
	})

	Context("SetParent", func() {
		It("should allow setting node parent", func() {
			anotherParent := &signalingfakes.FakeReceiverNode{}
			node.SetParent(anotherParent)
			Expect(node.GetParent()).To(Equal(anotherParent))
		})
	})
})
