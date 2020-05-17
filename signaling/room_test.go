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
			room *signaling.Room
		)

		It("should create new room", func() {
			room = signaling.NewRoom()
			Expect(room).ToNot(BeNil())
		})
	})

	Context("StartReaper", func() {
		var (
			room          *signaling.Room
			fakeMember    *signalingfakes.FakeRoomMember
			anotherMember *signalingfakes.FakeRoomMember
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

	Context("GetMember", func() {
		var (
			room          *signaling.Room
			fakeMember    *signalingfakes.FakeRoomMember
			anotherMember *signalingfakes.FakeRoomMember
		)

		BeforeEach(func() {
			room = signaling.NewRoom()

			fakeMember = &signalingfakes.FakeRoomMember{}
			fakeMember.IDReturns(signaling.PeerID("123"))
			room.AddMember(fakeMember)
		})

		It("should get one member", func() {
			Expect(room.GetMember(fakeMember.ID())).To(Equal(fakeMember))
		})

		It("should get two members", func() {
			anotherMember = &signalingfakes.FakeRoomMember{}
			anotherMember.IDReturns(signaling.PeerID("124"))
			room.AddMember(anotherMember)

			Expect(room.GetMember(fakeMember.ID())).To(Equal(fakeMember))
			Expect(room.GetMember(anotherMember.ID())).To(Equal(anotherMember))
		})

		It("should get two members with unique IDs", func() {
			anotherMember = &signalingfakes.FakeRoomMember{}
			anotherMember.IDReturns(signaling.PeerID("124"))
			room.AddMember(anotherMember)

			Expect(fakeMember.ID()).ToNot(Equal(anotherMember.ID()))
		})
	})

	Context("GetMemberCount", func() {
		var (
			room          *signaling.Room
			fakeMember    *signalingfakes.FakeRoomMember
			anotherMember *signalingfakes.FakeRoomMember
		)

		BeforeEach(func() {
			room = signaling.NewRoom()
			fakeMember = &signalingfakes.FakeRoomMember{}
			fakeMember.IDReturns(signaling.PeerID("123"))

			room.AddMember(fakeMember)
		})

		It("should get member count equal to one", func() {
			Expect(room.GetMemberCount()).To(Equal(1))
			Expect(room.GetMember(fakeMember.ID())).To(Equal(fakeMember))
		})

		It("should get member count equal to two", func() {
			anotherMember = &signalingfakes.FakeRoomMember{}
			anotherMember.IDReturns(signaling.PeerID("124"))
			room.AddMember(anotherMember)

			Expect(room.GetMemberCount()).To(Equal(2))
			Expect(room.GetMember(fakeMember.ID())).To(Equal(fakeMember))
		})

		It("should get member count equal to zero", func() {
			room.RemoveMember(fakeMember)
			Expect(room.GetMemberCount()).To(Equal(0))
		})
	})

	Context("AddMember", func() {
		var (
			room       *signaling.Room
			fakeMember *signalingfakes.FakeRoomMember
		)

		BeforeEach(func() {
			room = signaling.NewRoom()
			fakeMember = &signalingfakes.FakeRoomMember{}
			fakeMember.IDReturns(signaling.PeerID("123"))

			room.AddMember(fakeMember)
		})

		It("should add member", func() {
			Expect(room.GetMember(fakeMember.ID())).To(Equal(fakeMember))
			Expect(room.GetMemberCount()).To(Equal(1))
		})

		It("should not add existing member", func() {
			room.AddMember(fakeMember)

			Expect(room.GetMember(fakeMember.ID())).To(Equal(fakeMember))
			Expect(room.GetMemberCount()).To(Equal(1))
		})
	})

	Context("RemoveMember", func() {
		var (
			room          *signaling.Room
			fakeMember    *signalingfakes.FakeRoomMember
			anotherMember *signalingfakes.FakeRoomMember
		)

		BeforeEach(func() {
			room = signaling.NewRoom()

			fakeMember = &signalingfakes.FakeRoomMember{}
			fakeMember.IDReturns(signaling.PeerID("123"))

			room.AddMember(fakeMember)
		})

		It("should remove member", func() {
			Expect(room.GetMember(fakeMember.ID())).To(Equal(fakeMember))

			room.RemoveMember(fakeMember)
			Expect(room.GetMember(fakeMember.ID())).To(BeNil())
			Expect(room.GetMemberCount()).To(Equal(0))
		})

		It("should remove only one member", func() {
			anotherMember = &signalingfakes.FakeRoomMember{}
			anotherMember.IDReturns(signaling.PeerID("124"))

			room.AddMember(fakeMember)
			room.AddMember(anotherMember)

			Expect(room.GetMemberCount()).To(Equal(2))

			room.RemoveMember(fakeMember)
			Expect(room.GetMemberCount()).To(Equal(1))
			Expect(room.GetMember(fakeMember.ID())).To(BeNil())
			Expect(room.GetMember(anotherMember.ID())).ToNot(BeNil())
		})

		It("should remove only two members", func() {
			anotherMember = &signalingfakes.FakeRoomMember{}
			anotherMember.IDReturns(signaling.PeerID("124"))

			room.AddMember(anotherMember)

			Expect(room.GetMemberCount()).To(Equal(2))

			room.RemoveMember(fakeMember)
			Expect(room.GetMemberCount()).To(Equal(1))
			Expect(room.GetMember(fakeMember.ID())).To(BeNil())
			Expect(room.GetMember(anotherMember.ID())).ToNot(BeNil())
			Expect(room.GetMember(anotherMember.ID())).To(Equal(anotherMember))

			room.RemoveMember(anotherMember)
			Expect(room.GetMemberCount()).To(Equal(0))
			Expect(room.GetMember(anotherMember.ID())).To(BeNil())
			Expect(room.GetMember(anotherMember.ID())).To(BeNil())
		})
	})
})
