package signaling_test

import (
	"github.com/ryanrolds/club/pkg/signaling"
	"github.com/ryanrolds/club/pkg/signaling/signalingfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Room", func() {
	var (
		room          *signaling.Room
		defaultGroup  *signalingfakes.FakeReceiverGroup
		testGroup     *signalingfakes.FakeReceiverGroup
		fakeDependent *signalingfakes.FakeReceiverNode
	)

	BeforeEach(func() {
		room = signaling.NewRoom()

		defaultGroup = &signalingfakes.FakeReceiverGroup{}
		defaultGroup.IDReturns(signaling.DefaultGroupID)
		testGroup = &signalingfakes.FakeReceiverGroup{}
		testGroup.IDReturns("test")

		err := room.AddGroup(defaultGroup)
		Expect(err).To(BeNil())
		err = room.AddGroup(testGroup)
		Expect(err).To(BeNil())

		fakeDependent = &signalingfakes.FakeReceiverNode{}
		fakeDependent.IDReturns(signaling.NodeID("123"))

		room.AddDependent(fakeDependent)
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
				Expect(defaultGroup.AddDependentCallCount()).To(Equal(0))
				Expect(fakeDependent.SetParentCallCount()).To(Equal(0))

				room.Receive(signaling.NewJoinMessage(fakeDependent.ID()))

				Expect(defaultGroup.AddDependentCallCount()).To(Equal(1))
				newDependent := defaultGroup.AddDependentArgsForCall(0)
				Expect(newDependent).To(Equal(fakeDependent))

				Expect(fakeDependent.SetParentCallCount()).To(Equal(1))
				newParent := fakeDependent.SetParentArgsForCall(0)
				Expect(newParent).To(Equal(defaultGroup))

				Expect(testGroup.AddDependentCallCount()).To(Equal(0))
			})

			It("should allow adding to non-default group", func() {
				Expect(testGroup.AddDependentCallCount()).To(Equal(0))
				Expect(fakeDependent.SetParentCallCount()).To(Equal(0))

				message := signaling.NewJoinMessage(fakeDependent.ID())
				message.Payload = signaling.MessagePayload{
					signaling.MessagePayloadKeyGroup: "test",
				}

				room.Receive(message)

				Expect(testGroup.AddDependentCallCount()).To(Equal(1))
				newDependent := testGroup.AddDependentArgsForCall(0)
				Expect(newDependent).To(Equal(fakeDependent))

				Expect(fakeDependent.SetParentCallCount()).To(Equal(1))
				newParent := fakeDependent.SetParentArgsForCall(0)
				Expect(newParent).To(Equal(testGroup))

				Expect(defaultGroup.AddDependentCallCount()).To(Equal(0))
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
})
