package signaling_test

import (
	"github.com/ryanrolds/club/pkg/signaling"
	"github.com/ryanrolds/club/pkg/signaling/signalingfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server", func() {
	var fakeRoom *signalingfakes.FakeDispatcher

	BeforeEach(func() {
		fakeRoom = &signalingfakes.FakeDispatcher{}
	})

	Context("NewServer", func() {
		It("should return new server", func() {
			server := signaling.NewServer(fakeRoom)
			Expect(server).ToNot(BeNil())
		})
	})

	Context("ServeHTTP", func() {
		It("should dispatch message to room", func() {

		})

		It("should send error message to client if processing errs", func() {

		})
	})
})
