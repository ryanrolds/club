package signaling_test

import (
	"github.com/ryanrolds/club/pkg/signaling"
	"github.com/ryanrolds/club/pkg/signaling/signalingfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GroupNode", func() {
	var (
		group            signaling.GroupNode
		room             *signaling.Room
		fakeDependent    *signalingfakes.FakeReceiverNode
		anotherDependent *signalingfakes.FakeReceiverNode
	)

	BeforeEach(func() {
		room = &signaling.Room{}
		group = signaling.NewGroupNode("foo", room, 12)

		fakeDependent = &signalingfakes.FakeReceiverNode{}
		fakeDependent.IDReturns(signaling.NodeID("123"))
		group.AddDependent(fakeDependent)

		anotherDependent = &signalingfakes.FakeReceiverNode{}
		anotherDependent.IDReturns(signaling.NodeID("456"))
		group.AddDependent(anotherDependent)
	})

	Context("NewGroupNode", func() {
		It("should create new group", func() {
			group = signaling.NewGroupNode("id", room, 42)
			Expect(group).ToNot(BeNil())
		})
	})

	Context("Receive", func() {
		Context("Leave message", func() {
			It("should remove dependent", func() {
				Expect(fakeDependent.ReceiveCallCount()).To(Equal(1))

				Expect(group.GetDependent(anotherDependent.ID())).To(Equal(anotherDependent))
				group.Receive(signaling.NewLeaveMessage(anotherDependent.ID()))
				Expect(group.GetDependent(anotherDependent.ID())).To(BeNil())

				Expect(fakeDependent.ReceiveCallCount()).To(Equal(2))
				message := fakeDependent.ReceiveArgsForCall(1)
				Expect(message.Type).To(Equal(signaling.MessageTypeLeave))
				Expect(message.SourceID).To(Equal(anotherDependent.ID()))
			})

			It("should do nothing if dependent does not exist", func() {
				Expect(fakeDependent.ReceiveCallCount()).To(Equal(1))

				group.Receive(signaling.NewLeaveMessage(signaling.NodeID("doesnotexist")))

				Expect(fakeDependent.ReceiveCallCount()).To(Equal(1))
				Expect(anotherDependent.ReceiveCallCount()).To(Equal(0))
				Expect(group.GetDependent(fakeDependent.ID())).To(Equal(fakeDependent))
				Expect(group.GetDependent(anotherDependent.ID())).To(Equal(anotherDependent))
			})
		})

		Context("RTC related messages", func() {
			testRTCMessage := func(messsageType signaling.MessageType) {
				group.AddDependent(anotherDependent)

				Expect(fakeDependent.ReceiveCallCount()).To(Equal(1))
				Expect(anotherDependent.ReceiveCallCount()).To(Equal(0))

				group.Receive(signaling.Message{
					Type:          messsageType,
					SourceID:      fakeDependent.ID(),
					DestinationID: anotherDependent.ID(),
				})

				Expect(fakeDependent.ReceiveCallCount()).To(Equal(1))
				Expect(anotherDependent.ReceiveCallCount()).To(Equal(1))

				message := anotherDependent.ReceiveArgsForCall(0)
				Expect(message.Type).To(Equal(messsageType))
				Expect(message.SourceID).To(Equal(fakeDependent.ID()))
			}

			Context("MessageTypeOffer", func() {
				It("should send message to intended destination", func() {
					testRTCMessage(signaling.MessageTypeICECandidate)
				})
			})

			Context("MessageTypeAnswer", func() {
				It("should send message to intended destination", func() {
					testRTCMessage(signaling.MessageTypeAnswer)
				})
			})

			Context("MessageTypeICECandidate", func() {
				It("should send message to intended destination", func() {
					testRTCMessage(signaling.MessageTypeICECandidate)
				})
			})
		})
	})
})
