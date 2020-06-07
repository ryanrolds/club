package signaling_test

import (
	"github.com/ryanrolds/club/pkg/signaling"
	"github.com/ryanrolds/club/pkg/signaling/signalingfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Peer", func() {
	var (
		fakeConn   *signalingfakes.FakeWebsocketConn
		fakeParent *signalingfakes.FakeReceiverNode
		peer       *signaling.WebsocketPeer
		// validMessage = []byte(`{"type":"type","destId":"destID","payload":{}}`)
	)

	BeforeEach(func() {
		fakeConn = &signalingfakes.FakeWebsocketConn{}
		peer = signaling.NewWebsocketPeer(fakeConn, fakeParent)
	})

	Context("NewPeer", func() {
		It("should create new peer", func() {
			peer = signaling.NewWebsocketPeer(fakeConn, fakeParent)
			Expect(peer).ToNot(BeNil())
			Expect(peer.ID()).ToNot(Equal(signaling.PeerID("")))
		})
	})

	Context("Close", func() {
		It("should close connection", func() {
			peer.Close()
			Expect(fakeConn.CloseCallCount()).To(Equal(1))
		})
	})
})
