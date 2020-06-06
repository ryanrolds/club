package signaling_test

import (
	"github.com/ryanrolds/club/pkg/signaling"
	"github.com/ryanrolds/club/pkg/signaling/signalingfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Group", func() {
	var (
		group         *signaling.Group
		fakeMember    *signalingfakes.FakeRoomMember
		anotherMember *signalingfakes.FakeRoomMember
	)

	BeforeEach(func() {
		group = signaling.NewGroup("foo", 12)

		fakeMember = &signalingfakes.FakeRoomMember{}
		fakeMember.IDReturns(signaling.PeerID("123"))
		fakeMember.GetGroupReturns(group)
		group.AddMember(fakeMember)
	})

	Context("NewGroup", func() {
		It("should create new group", func() {
			group = signaling.NewGroup("id", 42)
			Expect(group).ToNot(BeNil())
		})
	})

	Context("ID", func() {
		It("should return ID", func() {
			Expect(group.ID()).To(Equal(signaling.GroupID("foo")))
		})
	})

	Context("GetMember", func() {
		It("should get one member", func() {
			Expect(group.GetMember(fakeMember.ID())).To(Equal(fakeMember))
		})

		It("should get two members", func() {
			anotherMember = &signalingfakes.FakeRoomMember{}
			anotherMember.IDReturns(signaling.PeerID("124"))
			group.AddMember(anotherMember)

			Expect(group.GetMember(fakeMember.ID())).To(Equal(fakeMember))
			Expect(group.GetMember(anotherMember.ID())).To(Equal(anotherMember))
		})

		It("should get two members with unique IDs", func() {
			anotherMember = &signalingfakes.FakeRoomMember{}
			anotherMember.IDReturns(signaling.PeerID("124"))
			group.AddMember(anotherMember)

			Expect(fakeMember.ID()).ToNot(Equal(anotherMember.ID()))
		})
	})

	Context("GetMemberCount", func() {
		It("should get member count equal to one", func() {
			Expect(group.GetMemberCount()).To(Equal(1))
			Expect(group.GetMember(fakeMember.ID())).To(Equal(fakeMember))
		})

		It("should get member count equal to two", func() {
			anotherMember = &signalingfakes.FakeRoomMember{}
			anotherMember.IDReturns(signaling.PeerID("124"))
			group.AddMember(anotherMember)

			Expect(group.GetMemberCount()).To(Equal(2))
			Expect(group.GetMember(fakeMember.ID())).To(Equal(fakeMember))
		})

		It("should get member count equal to zero", func() {
			group.RemoveMember(fakeMember)
			Expect(group.GetMemberCount()).To(Equal(0))
		})
	})

	Context("AddMember", func() {
		It("should add member", func() {
			Expect(group.GetMember(fakeMember.ID())).To(Equal(fakeMember))
			Expect(group.GetMemberCount()).To(Equal(1))
		})

		It("should not add existing member", func() {
			group.AddMember(fakeMember)

			Expect(group.GetMember(fakeMember.ID())).To(Equal(fakeMember))
			Expect(group.GetMemberCount()).To(Equal(1))
		})
	})

	Context("RemoveMember", func() {
		It("should remove member", func() {
			Expect(group.GetMember(fakeMember.ID())).To(Equal(fakeMember))

			group.RemoveMember(fakeMember)
			Expect(group.GetMember(fakeMember.ID())).To(BeNil())
			Expect(group.GetMemberCount()).To(Equal(0))
		})

		It("should remove only one member", func() {
			anotherMember = &signalingfakes.FakeRoomMember{}
			anotherMember.IDReturns(signaling.PeerID("124"))

			group.AddMember(fakeMember)
			group.AddMember(anotherMember)

			Expect(group.GetMemberCount()).To(Equal(2))

			group.RemoveMember(fakeMember)
			Expect(group.GetMemberCount()).To(Equal(1))
			Expect(group.GetMember(fakeMember.ID())).To(BeNil())
			Expect(group.GetMember(anotherMember.ID())).ToNot(BeNil())
		})

		It("should remove only two members", func() {
			anotherMember = &signalingfakes.FakeRoomMember{}
			anotherMember.IDReturns(signaling.PeerID("124"))

			group.AddMember(anotherMember)

			Expect(group.GetMemberCount()).To(Equal(2))

			group.RemoveMember(fakeMember)
			Expect(group.GetMemberCount()).To(Equal(1))
			Expect(group.GetMember(fakeMember.ID())).To(BeNil())
			Expect(group.GetMember(anotherMember.ID())).ToNot(BeNil())
			Expect(group.GetMember(anotherMember.ID())).To(Equal(anotherMember))

			group.RemoveMember(anotherMember)
			Expect(group.GetMemberCount()).To(Equal(0))
			Expect(group.GetMember(anotherMember.ID())).To(BeNil())
			Expect(group.GetMember(anotherMember.ID())).To(BeNil())
		})
	})

	Context("MessageMember", func() {
		It("should message member", func() {

		})
	})

	Context("Broadcast", func() {
		It("should message all mambers except source", func() {

		})
	})
})
