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
	})

	Context("NewRoom", func() {
		It("should create new room", func() {
			room = signaling.NewRoom()
			Expect(room).ToNot(BeNil())
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
