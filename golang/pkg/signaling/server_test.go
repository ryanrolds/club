package signaling_test


import (
	"github.com/ryanrolds/club/pkg/signaling"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


var _ = Describe("Server", func() {
  Context("NewServer", func() {
    It("should return new server", func() {

    })
  })

  Context("ServeHTTP", func() {
    It("should dispatch message to room", func() {

    })

    It("should send error message to client if processing errs", func() {

    })
  })
})
