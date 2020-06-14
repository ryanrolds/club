package signaling_test

import (
	"github.com/ryanrolds/club/pkg/signaling"
	"github.com/ryanrolds/club/pkg/signaling/signalingfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GroupNode", func() {
	var (
		group                     signaling.GroupNode
		room                      *signaling.Room
		fakeMember                *signalingfakes.FakeReceiverNode
		anotherMember             *signalingfakes.FakeReceiverNode
		fakeMemberReceiveCount    int
		anotherMemberReceiveCount int
	)

	BeforeEach(func() {
		room = &signaling.Room{}
		group = signaling.NewGroupNode("foo", room, 12)

		fakeMember = &signalingfakes.FakeReceiverNode{}
		fakeMember.IDReturns(signaling.NodeID("123"))
		group.AddMember(fakeMember)

		anotherMember = &signalingfakes.FakeReceiverNode{}
		anotherMember.IDReturns(signaling.NodeID("456"))
		group.AddMember(anotherMember)

		fakeMemberReceiveCount = fakeMember.ReceiveCallCount()
		Expect(fakeMemberReceiveCount).To(Equal(2))
		anotherMemberReceiveCount = anotherMember.ReceiveCallCount()
		Expect(anotherMemberReceiveCount).To(Equal(1))
	})

	Context("NewGroupNode", func() {
		It("should create new group", func() {
			group = signaling.NewGroupNode("id", room, 42)
			Expect(group).ToNot(BeNil())
		})
	})

	Context("Receive", func() {
		Context("Leave message", func() {
			It("should remove member", func() {
				Expect(group.GetMember(anotherMember.ID())).To(Equal(anotherMember))
				group.Receive(signaling.NewLeaveMessage(anotherMember.ID()))
				Expect(group.GetMember(anotherMember.ID())).To(BeNil())

				Expect(fakeMember.ReceiveCallCount()).To(Equal(fakeMemberReceiveCount + 1))
				message := fakeMember.ReceiveArgsForCall(fakeMemberReceiveCount)
				Expect(message.Type).To(Equal(signaling.MessageTypeLeave))
				Expect(message.SourceID).To(Equal(anotherMember.ID()))
			})

			It("should do nothing if member does not exist", func() {
				group.Receive(signaling.NewLeaveMessage(signaling.NodeID("doesnotexist")))

				Expect(fakeMember.ReceiveCallCount()).To(Equal(fakeMemberReceiveCount))
				Expect(anotherMember.ReceiveCallCount()).To(Equal(anotherMemberReceiveCount))
				Expect(group.GetMember(fakeMember.ID())).To(Equal(fakeMember))
				Expect(group.GetMember(anotherMember.ID())).To(Equal(anotherMember))
			})
		})

		Context("RTC related messages", func() {
			testRTCMessage := func(messsageType signaling.MessageType) {
				group.Receive(signaling.Message{
					Type:          messsageType,
					SourceID:      fakeMember.ID(),
					DestinationID: anotherMember.ID(),
				})

				Expect(fakeMember.ReceiveCallCount()).To(Equal(fakeMemberReceiveCount))
				Expect(anotherMember.ReceiveCallCount()).To(Equal(anotherMemberReceiveCount + 1))

				message := anotherMember.ReceiveArgsForCall(anotherMemberReceiveCount)
				Expect(message.Type).To(Equal(messsageType))
				Expect(message.SourceID).To(Equal(fakeMember.ID()))
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

	Context("GetDetails", func() {
		It("should return the group details", func() {
			details := group.GetDetails()
			Expect(details.ID).To(Equal(signaling.NodeID("foo")))
			Expect(details.Name).To(Equal("foo"))
			Expect(details.Limit).To(Equal(12))
			Expect(details.MemberCount).To(Equal(2))
		})
	})
})
