package signaling_test

import (
	"time"

	"github.com/ryanrolds/club/pkg/signaling"
	"github.com/ryanrolds/club/pkg/signaling/signalingfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Room", func() {
	Context("NewServer", func() {
		It("should create new server", func() {
			room := signaling.NewRoom()
			Expect(room).ToNot(BeNil())
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
