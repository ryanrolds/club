package signaling_test

import (
	"github.com/ryanrolds/club/pkg/signaling"
	"github.com/ryanrolds/club/pkg/signaling/signalingfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Peer", func() {
	var (
		fakeConn     *signalingfakes.FakePeerConnection
		peer         *signaling.Peer
		validMessage = []byte(`{"type":"type","destId":"destID","payload":{}}`)
	)

	BeforeEach(func() {
		fakeConn = &signalingfakes.FakePeerConnection{}
		peer = signaling.NewPeer(fakeConn)
	})

	Context("NewPeer", func() {
		It("should create new peer", func() {
			peer = signaling.NewPeer(fakeConn)
			Expect(peer).ToNot(BeNil())
			Expect(peer.ID()).ToNot(Equal(signaling.PeerID("")))
		})
	})

	Context("GetNextMessage", func() {
		It("should return next message", func() {
			fakeConn.ReadMessageReturns(1, validMessage, nil)

			message, err := peer.GetNextMessage()
			Expect(err).To(BeNil())
			Expect(message.Type).To(Equal(signaling.MessageType("type")))
			Expect(message.SourceID).To(Equal(peer.ID()))
		})
	})

	Context("SendMessage", func() {
		It("should send JSON message", func() {
			message := signaling.Message{
				Type:          "join",
				SourceID:      peer.ID(),
				DestinationID: signaling.PeerID("destination"),
				Payload: map[string]interface{}{
					"foo": "bar",
				},
			}
			err := peer.SendMessage(message)
			Expect(err).To(BeNil())

			sentMsg := fakeConn.WriteJSONArgsForCall(0)
			Expect(sentMsg).To(Equal(message))
		})
	})

	Context("Close", func() {
		It("should close connection", func() {
			peer.Close()
			Expect(fakeConn.CloseCallCount()).To(Equal(1))
		})
	})
})
