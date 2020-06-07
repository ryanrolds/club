package signaling_test

import (
	"github.com/ryanrolds/club/pkg/signaling"
	"github.com/ryanrolds/club/pkg/signaling/signalingfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GroupNode", func() {
	var (
		group         signaling.GroupNode
		room          *signaling.Room
		fakeDependent *signalingfakes.FakeReceiverNode
		// anotherDependent *signalingfakes.FakeReceiverNode
	)

	BeforeEach(func() {
		room = &signaling.Room{}
		group = signaling.NewGroupNode("foo", room, 12)

		fakeDependent = &signalingfakes.FakeReceiverNode{}
		fakeDependent.IDReturns(signaling.NodeID("123"))
		group.AddDependent(fakeDependent)
	})

	Context("NewGroupNode", func() {
		It("should create new group", func() {
			group = signaling.NewGroupNode("id", room, 42)
			Expect(group).ToNot(BeNil())
		})
	})

	Context("Receive", func() {
		Context("Leave message", func() {

		})

		Context("RTC related messages", func() {

		})
	})
})
