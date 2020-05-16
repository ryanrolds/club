package signaling_test

import (
	"time"

	"github.com/ryanrolds/club/signaling"
	"github.com/ryanrolds/club/signaling/signalingfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Room", func() {
	Context("NewRoom", func() {
		var (
			room             *signaling.Room
			fakeMember       *signalingfakes.FakeRoomMember
			fakeSecondMember *signalingfakes.FakeRoomMember
		)

		BeforeEach(func() {
			fakeMember = &signalingfakes.FakeRoomMember{}
			fakeMember.IDReturns(signaling.PeerID("123"))

			fakeSecondMember = &signalingfakes.FakeRoomMember{}
			fakeSecondMember.IDReturns(signaling.PeerID("124"))

		})

		It("should create new room", func() {
			room = signaling.NewRoom()
			Expect(room).ToNot(BeNil())
		})

		It("should add member", func() {
			room = signaling.NewRoom()

			room.AddMember(fakeMember)
			Expect(room.GetMemberCount()).To(Equal(1))
		})

		It("should not add existing member", func() {
			room = signaling.NewRoom()

			room.AddMember(fakeMember)
			room.AddMember(fakeMember)

			Expect(room.GetMemberCount()).To(Equal(1))
		})

		It("should guarantee unique member IDs", func() {
			room = signaling.NewRoom()

			room.AddMember(fakeMember)
			room.AddMember(fakeSecondMember)

			Expect(fakeMember.ID()).ToNot(Equal(fakeSecondMember.ID()))
		})

		It("Should get member", func() {
			room = signaling.NewRoom()

			room.AddMember(fakeMember)
			Expect(room.GetMember(fakeMember.ID())).To(Equal(fakeMember))
		})

		It("should remove member", func() {
			room = signaling.NewRoom()

			room.AddMember(fakeMember)
			Expect(room.GetMember(fakeMember.ID())).To(Equal(fakeMember))

			room.RemoveMember(fakeMember)
			Expect(room.GetMember(fakeMember.ID())).To(BeNil())
			Expect(room.GetMemberCount()).To(Equal(0))
		})

		It("should remove only one member", func() {
			room = signaling.NewRoom()

			room.AddMember(fakeMember)
			room.AddMember(fakeSecondMember)
			Expect(room.GetMemberCount()).To(Equal(2))

			room.RemoveMember(fakeMember)
			Expect(room.GetMemberCount()).To(Equal(1))
			Expect(room.GetMember(fakeMember.ID())).To(BeNil())
			Expect(room.GetMember(fakeSecondMember.ID())).ToNot(BeNil())
		})
	})

	Context("StartReaper", func() {
		var (
			fakeMember    *signalingfakes.FakeRoomMember
			anotherMember *signalingfakes.FakeRoomMember
			room          *signaling.Room
		)

		BeforeEach(func() {
			room = signaling.NewRoom()

			fakeMember = &signalingfakes.FakeRoomMember{}
			fakeMember.IDReturns(signaling.PeerID("123"))
			room.AddMember(fakeMember)

			anotherMember = &signalingfakes.FakeRoomMember{}
			anotherMember.IDReturns(signaling.PeerID("42"))
			room.AddMember(anotherMember)
		})

		It("should run every interval", func() {
			room.StartReaper(time.Millisecond * 250)
			time.Sleep(time.Second)
			Expect(fakeMember.TimedoutCallCount()).To(Equal(4))
			Expect(anotherMember.TimedoutCallCount()).To(Equal(4))
		})

		It("should remove member if timedout and leave non-timedout members", func() {
			fakeMember.TimedoutReturns(true)
			anotherMember.TimedoutReturns(false)

			Expect(room.GetMember(fakeMember.ID())).To(Equal(fakeMember))
			Expect(room.GetMember(anotherMember.ID())).To(Equal(anotherMember))

			room.StartReaper(time.Millisecond * 50)

			time.Sleep(time.Millisecond * 100)

			Expect(room.GetMember(fakeMember.ID())).To(BeNil())
			Expect(room.GetMember(anotherMember.ID())).To(Equal(anotherMember))
		})

		It("should inform other members", func() {
			fakeMember.TimedoutReturns(true)

			room.StartReaper(time.Millisecond * 75)

			time.Sleep(time.Millisecond * 100)

			Expect(fakeMember.SendMessageCallCount()).To(Equal(0))
			Expect(anotherMember.SendMessageCallCount()).To(Equal(1))

			message := anotherMember.SendMessageArgsForCall(0)
			Expect(message.Type).To(Equal(signaling.MessageTypeLeave))
		})
	})
})
