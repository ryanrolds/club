package signaling_test

import (
	"github.com/ryanrolds/club/pkg/signaling"
	"github.com/ryanrolds/club/pkg/signaling/signalingfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Room", func() {
	var (
		room         *signaling.Room
		defaultGroup *signalingfakes.FakeReceiverGroup
		testGroup    *signalingfakes.FakeReceiverGroup
		fakeMember   *signalingfakes.FakeReceiverNode
	)

	BeforeEach(func() {
		room = signaling.NewRoom()

		defaultGroup = &signalingfakes.FakeReceiverGroup{}
		defaultGroup.IDReturns(signaling.DefaultGroupID)
		defaultGroup.GetDetailsReturns(signaling.GroupDetails{
			ID:          signaling.DefaultGroupID,
			Name:        "foo",
			Limit:       42,
			MemberCount: 2,
		})

		testGroup = &signalingfakes.FakeReceiverGroup{}
		testGroup.IDReturns("test")
		testGroup.GetDetailsReturns(signaling.GroupDetails{
			ID:          signaling.DefaultGroupID,
			Name:        "foo",
			Limit:       42,
			MemberCount: 2,
		})

		err := room.AddGroup(defaultGroup)
		Expect(err).To(BeNil())
		err = room.AddGroup(testGroup)
		Expect(err).To(BeNil())

		fakeMember = &signalingfakes.FakeReceiverNode{}
		fakeMember.IDReturns(signaling.NodeID("123"))

		room.AddMember(fakeMember)
	})

	Context("NewRoom", func() {
		It("should create new room", func() {
			room = signaling.NewRoom()
			Expect(room).ToNot(BeNil())
		})
	})

	Context("Receive", func() {
		Context("MessageTypeJoin", func() {
			It("should add the depenent to the group", func() {
				Expect(defaultGroup.AddMemberCallCount()).To(Equal(0))
				Expect(fakeMember.SetParentCallCount()).To(Equal(0))

				room.Receive(signaling.NewJoinMessage(fakeMember.ID()))

				Expect(defaultGroup.AddMemberCallCount()).To(Equal(1))
				newMember := defaultGroup.AddMemberArgsForCall(0)
				Expect(newMember).To(Equal(fakeMember))

				Expect(fakeMember.SetParentCallCount()).To(Equal(1))
				newParent := fakeMember.SetParentArgsForCall(0)
				Expect(newParent).To(Equal(defaultGroup))

				Expect(testGroup.AddMemberCallCount()).To(Equal(0))
			})

			It("should allow adding to non-default group", func() {
				Expect(testGroup.AddMemberCallCount()).To(Equal(0))
				Expect(fakeMember.SetParentCallCount()).To(Equal(0))

				message := signaling.NewJoinMessage(fakeMember.ID())
				message.Payload = signaling.MessagePayload{
					signaling.MessagePayloadKeyGroup: "test",
				}

				room.Receive(message)

				Expect(testGroup.AddMemberCallCount()).To(Equal(1))
				newMember := testGroup.AddMemberArgsForCall(0)
				Expect(newMember).To(Equal(fakeMember))

				Expect(fakeMember.SetParentCallCount()).To(Equal(1))
				newParent := fakeMember.SetParentArgsForCall(0)
				Expect(newParent).To(Equal(testGroup))

				Expect(defaultGroup.AddMemberCallCount()).To(Equal(0))
			})
		})
	})

	Context("AddGroup", func() {
		It("should add group to room", func() {
			foo := &signalingfakes.FakeReceiverGroup{}
			foo.IDReturns("foo")

			err := room.AddGroup(foo)
			Expect(err).To(BeNil())

			Expect(room.GetGroup("foo")).To(Equal(foo))
		})

		It("should error if group is nil", func() {
			err := room.AddGroup(nil)
			Expect(err).To(Equal(signaling.ErrNonNilGroupRequired))
		})

		It("should error if group is already added to room", func() {
			err := room.AddGroup(defaultGroup)
			Expect(err).To(Equal(signaling.ErrGroupAlreadyExists))
		})
	})

	Context("GetGroup", func() {
		It("should return group if it exists", func() {
			Expect(room.GetGroup(signaling.DefaultGroupID)).To(Equal(defaultGroup))
		})

		It("should return nil if group does not exist", func() {
			Expect(room.GetGroup(signaling.NodeID("doesnotexist"))).To(BeNil())
		})
	})

	Context("GetDetailsForGroups", func() {
		It("should return array of group details", func() {
			details := room.GetDetailsForGroups()
			Expect(len(details)).To(Equal(2))
			Expect(details[0]).To(Equal(defaultGroup.GetDetails()))
			Expect(details[1]).To(Equal(testGroup.GetDetails()))
		})
	})
})
