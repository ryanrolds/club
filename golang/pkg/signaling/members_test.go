package signaling_test

import (
	"github.com/ryanrolds/club/pkg/signaling"
	"github.com/ryanrolds/club/pkg/signaling/signalingfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Members", func() {
	var (
		members       signaling.Members
		fakeParent    *signalingfakes.FakeReceiverNode
		fakeMember    *signalingfakes.FakeReceiverNode
		anotherMember *signalingfakes.FakeReceiverNode
	)

	BeforeEach(func() {
		members = signaling.NewMembers(12)

		fakeParent = &signalingfakes.FakeReceiverNode{}
		fakeParent.IDReturns(signaling.NodeID("parent"))

		fakeMember = &signalingfakes.FakeReceiverNode{}
		fakeMember.IDReturns(signaling.NodeID("123"))
		fakeMember.GetParentReturns(fakeParent)
		members.AddMember(fakeMember)

		anotherMember = &signalingfakes.FakeReceiverNode{}
		anotherMember.IDReturns(signaling.NodeID("456"))
		anotherMember.GetParentReturns(fakeParent)
	})

	Context("NewGroup", func() {
		It("should create new set of members", func() {
			members = signaling.NewMembers(42)
			Expect(members).ToNot(BeNil())
		})
	})

	Context("GetMember", func() {
		It("should get one member", func() {
			Expect(members.GetMember(fakeMember.ID())).To(Equal(fakeMember))
		})

		It("should get two members", func() {
			members.AddMember(anotherMember)

			Expect(members.GetMember(fakeMember.ID())).To(Equal(fakeMember))
			Expect(members.GetMember(anotherMember.ID())).To(Equal(anotherMember))
		})

		It("should return nil of member does not exist", func() {
			Expect(members.GetMember(signaling.NodeID("doesnotexist"))).To(BeNil())
		})
	})

	Context("GetLimit", func() {
		It("should return the limit of the set", func() {
			members = signaling.NewMembers(42)
			Expect(members.GetLimit()).To(Equal(42))
		})
	})

	Context("GetMembersCount", func() {
		It("should get member count equal to one", func() {
			Expect(members.GetMembersCount()).To(Equal(1))
			Expect(members.GetMember(fakeMember.ID())).To(Equal(fakeMember))
		})

		It("should get member count equal to two", func() {
			members.AddMember(anotherMember)

			Expect(members.GetMembersCount()).To(Equal(2))
			Expect(members.GetMember(fakeMember.ID())).To(Equal(fakeMember))
		})

		It("should get member count equal to zero", func() {
			members.RemoveMember(fakeMember)
			Expect(members.GetMembersCount()).To(Equal(0))
		})
	})

	Context("AddMember", func() {
		It("should add member", func() {
			Expect(members.GetMember(fakeMember.ID())).To(Equal(fakeMember))
			Expect(members.GetMembersCount()).To(Equal(1))
		})

		It("should not add existing member", func() {
			members.AddMember(fakeMember)

			Expect(members.GetMember(fakeMember.ID())).To(Equal(fakeMember))
			Expect(members.GetMembersCount()).To(Equal(1))
		})

		It("should inform other members of addition", func() {
			members.AddMember(anotherMember)

			Expect(fakeMember.ReceiveCallCount()).To(Equal(1))
			message := fakeMember.ReceiveArgsForCall(0)
			Expect(message.Type).To(Equal(signaling.MessageTypeJoin))
			Expect(message.SourceID).To(Equal(anotherMember.ID()))
		})
	})

	Context("RemoveMember", func() {
		It("should remove member", func() {
			Expect(members.GetMember(fakeMember.ID())).To(Equal(fakeMember))

			members.RemoveMember(fakeMember)
			Expect(members.GetMember(fakeMember.ID())).To(BeNil())
			Expect(members.GetMembersCount()).To(Equal(0))
		})

		It("should remove only one member", func() {
			members.AddMember(anotherMember)

			Expect(members.GetMembersCount()).To(Equal(2))

			members.RemoveMember(fakeMember)
			Expect(members.GetMembersCount()).To(Equal(1))
			Expect(members.GetMember(fakeMember.ID())).To(BeNil())
			Expect(members.GetMember(anotherMember.ID())).ToNot(BeNil())
		})

		It("should remove only two members", func() {
			anotherMember = &signalingfakes.FakeReceiverNode{}
			anotherMember.IDReturns(signaling.NodeID("124"))

			members.AddMember(anotherMember)

			Expect(members.GetMembersCount()).To(Equal(2))

			members.RemoveMember(fakeMember)
			Expect(members.GetMembersCount()).To(Equal(1))
			Expect(members.GetMember(fakeMember.ID())).To(BeNil())
			Expect(members.GetMember(anotherMember.ID())).ToNot(BeNil())
			Expect(members.GetMember(anotherMember.ID())).To(Equal(anotherMember))

			members.RemoveMember(anotherMember)
			Expect(members.GetMembersCount()).To(Equal(0))
			Expect(members.GetMember(anotherMember.ID())).To(BeNil())
			Expect(members.GetMember(anotherMember.ID())).To(BeNil())
		})

		It("should inform other members of removal", func() {
			members.AddMember(anotherMember)

			members.RemoveMember(fakeMember)

			Expect(anotherMember.ReceiveCallCount()).To(Equal(1))
			message := anotherMember.ReceiveArgsForCall(0)
			Expect(message.Type).To(Equal(signaling.MessageTypeLeave))
			Expect(message.SourceID).To(Equal(fakeMember.ID()))
		})
	})

	Context("MessageDependant", func() {
		It("should message dependant", func() {
			Expect(fakeMember.ReceiveCallCount()).To(Equal(0))

			members.AddMember(anotherMember)

			members.MessageMember(signaling.Message{
				Type:          signaling.MessageTypeJoin,
				SourceID:      signaling.NodeID("123"),
				DestinationID: signaling.NodeID("456"),
			})

			// Adding a member will call Receive on existing members
			Expect(fakeMember.ReceiveCallCount()).To(Equal(1))
			Expect(anotherMember.ReceiveCallCount()).To(Equal(1))
		})

		It("should handle trying to message a dependant that does not exist", func() {
			members.AddMember(anotherMember)

			members.MessageMember(signaling.Message{
				Type:          signaling.MessageTypeJoin,
				SourceID:      signaling.NodeID("123"),
				DestinationID: signaling.NodeID("doesnotexist"),
			})

			// Adding a member will call Receive on existing members
			Expect(fakeMember.ReceiveCallCount()).To(Equal(1))
			Expect(anotherMember.ReceiveCallCount()).To(Equal(0))
		})
	})

	Context("Broadcast", func() {
		It("should message all members except source", func() {
			members.AddMember(anotherMember)

			members.Broadcast(signaling.Message{
				Type:     signaling.MessageTypeJoin,
				SourceID: signaling.NodeID("abc"),
			})

			// Adding a member will call Receive on existing members
			Expect(fakeMember.ReceiveCallCount()).To(Equal(2))
			Expect(anotherMember.ReceiveCallCount()).To(Equal(1))
		})
	})
})
