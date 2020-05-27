package signaling_test

import (
	"time"

	"github.com/ryanrolds/club/pkg/signaling"
	"github.com/ryanrolds/club/pkg/signaling/signalingfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Room", func() {
	var (
		room         *signaling.Room
		defaultGroup *signaling.Group
		testGroup    *signaling.Group
	)

	BeforeEach(func() {
		room = signaling.NewRoom()

		defaultGroup = signaling.NewGroup(signaling.RoomDefaultGroupID, 12)
		testGroup = signaling.NewGroup("test", 42)

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

	Context("StartReaper", func() {
		It("should call prune on each group", func() {
			room = signaling.NewRoom()

			fakeGroup := &signalingfakes.FakeRoomGroup{}
			err := room.AddGroup(fakeGroup)
			Expect(err).To(BeNil())

			room.StartReaper(time.Millisecond * 50)

			time.Sleep(time.Millisecond * 75)

			callCount := fakeGroup.PruneStaleMembersCallCount()
			Expect(callCount).To(BeNumerically(">", 1))
		})
	})

	Context("Dispatch", func() {
		var (
			conn   *signalingfakes.FakePeerConnection
			member *signaling.Peer
		)

		BeforeEach(func() {
			conn = &signalingfakes.FakePeerConnection{}
			member = signaling.NewPeer(conn)
		})

		It("should add joining member to default group", func() {
			Expect(defaultGroup.GetMemberCount()).To(Equal(0))

			message := signaling.Message{
				Type:     signaling.MessageTypeJoin,
				SourceID: member.ID(),
				Payload:  signaling.MessagePayload{},
			}

			err := room.Dispatch(member, message)
			Expect(err).To(BeNil())

			Expect(defaultGroup.GetMemberCount()).To(Equal(1))
			Expect(defaultGroup.GetMember(member.ID())).To(Equal(member))
		})

		It("should add joining member to requested group", func() {
			Expect(testGroup.GetMemberCount()).To(Equal(0))

			message := signaling.Message{
				Type:     signaling.MessageTypeJoin,
				SourceID: member.ID(),
				Payload: signaling.MessagePayload{
					signaling.MessagePayloadKeyGroup: "test",
				},
			}

			err := room.Dispatch(member, message)
			Expect(err).To(BeNil())

			Expect(testGroup.GetMemberCount()).To(Equal(1))
			Expect(testGroup.GetMember(member.ID())).To(Equal(member))
		})

		It("should error if group does not exist", func() {
			Expect(testGroup.GetMemberCount()).To(Equal(0))

			message := signaling.Message{
				Type:     signaling.MessageTypeJoin,
				SourceID: member.ID(),
				Payload: signaling.MessagePayload{
					signaling.MessagePayloadKeyGroup: "doesnotexist",
				},
			}

			err := room.Dispatch(member, message)
			Expect(err).To(Equal(signaling.ErrGroupNotFound))
		})

		It("should process leave message", func() {
			defaultGroup.AddMember(member)
			member.SetGroup(defaultGroup)
			Expect(defaultGroup.GetMemberCount()).To(Equal(1))

			message := signaling.Message{
				Type:     signaling.MessageTypeLeave,
				SourceID: member.ID(),
				Payload:  signaling.MessagePayload{},
			}

			err := room.Dispatch(member, message)
			Expect(err).To(BeNil())

			Expect(defaultGroup.GetMemberCount()).To(Equal(0))
		})
	})

	Context("AddGroup", func() {
		It("should add group to room", func() {
			foo := signaling.NewGroup("foo", 10)

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
			Expect(room.GetGroup(signaling.RoomDefaultGroupID)).To(Equal(defaultGroup))
		})

		It("should return nil if group does not exist", func() {
			Expect(room.GetGroup(signaling.GroupID("doesnotexist"))).To(BeNil())
		})
	})
})
