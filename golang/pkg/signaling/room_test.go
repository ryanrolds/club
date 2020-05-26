package signaling_test

import (
	"github.com/ryanrolds/club/pkg/signaling"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Room", func() {
	var (
		room   *signaling.Room
		groupA *signaling.Group
		groupB *signaling.Group
	)

	BeforeEach(func() {
		room = signaling.NewRoom()

		groupA = signaling.NewGroup("groupA", 12)
		groupB = signaling.NewGroup("groupB", 42)

		err := room.AddGroup(groupA)
		Expect(err).To(BeNil())
		err = room.AddGroup(groupB)
		Expect(err).To(BeNil())
	})

	Context("NewRoom", func() {
		It("should create new room", func() {
			room = signaling.NewRoom()
			Expect(room).ToNot(BeNil())
		})
	})

	Context("AddGroup", func() {
		It("should add group to room", func() {

		})

		It("should error if group is nil", func() {

		})

		It("should error if group is already added to room", func() {

		})
	})

	Context("GetGroup", func() {
		It("should return group if it exists", func() {

		})

		It("should return nil if group does not exist", func() {

		})
	})
})
