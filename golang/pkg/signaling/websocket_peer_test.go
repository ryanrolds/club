package signaling_test

import (
	"errors"

	"github.com/ryanrolds/club/pkg/signaling"
	"github.com/ryanrolds/club/pkg/signaling/signalingfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Peer", func() {
	var (
		fakeConn        *signalingfakes.FakeWebsocketConn
		fakeParent      *signalingfakes.FakeReceiverNode
		peer            *signaling.WebsocketPeer
		shutdownMessage = signaling.Message{
			Type: signaling.MessageTypeShutdown,
		}
		kickMessage = signaling.Message{
			Type: signaling.MessageTypeKick,
		}
		validMessage = []byte(`{"type":"type","destId":"destID","payload":{}}`)
	)

	BeforeEach(func() {
		fakeConn = &signalingfakes.FakeWebsocketConn{}
		fakeParent = &signalingfakes.FakeReceiverNode{}
		peer = signaling.NewWebsocketPeer(fakeConn, fakeParent)
	})

	Context("NewPeer", func() {
		It("should create new peer", func() {
			peer = signaling.NewWebsocketPeer(fakeConn, fakeParent)
			Expect(peer).ToNot(BeNil())
			Expect(peer.ID()).ToNot(Equal(signaling.PeerID("")))
		})
	})

	Context("Receive", func() {
		It("should push message on channel", func() {
			peer.PumpWrite()

			message := signaling.Message{}
			peer.Receive(message)

			peer.Receive(shutdownMessage)

			peer.WaitForDisconnect()

			sentMessage := fakeConn.WriteJSONArgsForCall(0)
			Expect(sentMessage).To(Equal(message))
			sentMessage = fakeConn.WriteJSONArgsForCall(1)
			Expect(sentMessage).To(Equal(shutdownMessage))
		})
	})

	Context("PumpWrite", func() {
		It("should get message and write to client", func() {
			peer.PumpWrite()

			message := signaling.Message{}
			peer.Receive(message)

			peer.Receive(shutdownMessage)

			peer.WaitForDisconnect()

			sentMessage := fakeConn.WriteJSONArgsForCall(0)
			Expect(sentMessage).To(Equal(message))
			sentMessage = fakeConn.WriteJSONArgsForCall(1)
			Expect(sentMessage).To(Equal(shutdownMessage))
		})

		It("should close channel and end when it gets a kick message", func() {
			peer.PumpWrite()

			peer.Receive(kickMessage)

			peer.WaitForDisconnect()
		})

		It("should close channel and end when it gets a shutdown message", func() {
			peer.PumpWrite()

			peer.Receive(shutdownMessage)

			peer.WaitForDisconnect()
		})
	})

	Context("PumpRead", func() {
		It("should read message from client", func() {
			callCount := 0
			fakeConn.ReadMessageCalls(func() (int, []byte, error) {
				callCount++

				if callCount <= 1 {
					return 1, validMessage, nil
				}

				return 1, []byte{}, errors.New("something went wrong")
			})

			peer.PumpRead()

			peer.WaitForDisconnect()

			Expect(fakeParent.ReceiveCallCount()).To(Equal(2))
			message := fakeParent.ReceiveArgsForCall(0)
			Expect(message.Type).To(Equal(signaling.MessageType("type")))
			message = fakeParent.ReceiveArgsForCall(1)
			Expect(message.Type).To(Equal(signaling.MessageTypeLeave))
		})

		It("should not send heartbeats to parent", func() {
			callCount := 0
			fakeConn.ReadMessageCalls(func() (int, []byte, error) {
				callCount++

				if callCount <= 1 {
					return 1, []byte(`{"type":"heartbeat","destId":"destID","payload":{}}`), nil
				}

				return 1, []byte{}, errors.New("something went wrong")
			})

			peer.PumpRead()

			peer.WaitForDisconnect()

			Expect(fakeConn.ReadMessageCallCount()).To(Equal(2))
			Expect(fakeParent.ReceiveCallCount()).To(Equal(1))
			message := fakeParent.ReceiveArgsForCall(0)
			Expect(message.Type).To(Equal(signaling.MessageTypeLeave))
		})
	})
})
