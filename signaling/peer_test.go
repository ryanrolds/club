package signaling_test

import (
	"github.com/gorilla/websocket"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ryanrolds/club/signaling"
)

var _ = Describe("Peer", func() {
	Context("NewPeer", func() {
		It("should create new peer", func() {
			conn := &websocket.Conn{}
			peer := signaling.NewPeer(conn)
			Expect(peer).ToNot(BeNil())
			Expect(peer.ID).ToNot(Equal(signaling.PeerID("")))
		})
	})

	Context("GetNextMessage", func() {

	})

	Context("SendMessage", func() {

	})

	Context("Leave", func() {

	})
})
