package signaling_test

import (
	"github.com/ryanrolds/club/pkg/signaling"
	"github.com/ryanrolds/club/pkg/signaling/signalingfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dependents", func() {
	var (
		dependents       signaling.Dependents
		fakeParent       *signalingfakes.FakeReceiverNode
		fakeDependent    *signalingfakes.FakeReceiverNode
		anotherDependent *signalingfakes.FakeReceiverNode
	)

	BeforeEach(func() {
		dependents = signaling.NewDependents(12)

		fakeParent = &signalingfakes.FakeReceiverNode{}
		fakeParent.IDReturns(signaling.NodeID("parent"))

		fakeDependent = &signalingfakes.FakeReceiverNode{}
		fakeDependent.IDReturns(signaling.NodeID("123"))
		fakeDependent.GetParentReturns(fakeParent)
		dependents.AddDependent(fakeDependent)

		anotherDependent = &signalingfakes.FakeReceiverNode{}
		anotherDependent.IDReturns(signaling.NodeID("456"))
		anotherDependent.GetParentReturns(fakeParent)
	})

	Context("NewGroup", func() {
		It("should create new set of dependents", func() {
			dependents = signaling.NewDependents(42)
			Expect(dependents).ToNot(BeNil())
		})
	})

	Context("GetDependent", func() {
		It("should get one dependent", func() {
			Expect(dependents.GetDependent(fakeDependent.ID())).To(Equal(fakeDependent))
		})

		It("should get two dependents", func() {
			dependents.AddDependent(anotherDependent)

			Expect(dependents.GetDependent(fakeDependent.ID())).To(Equal(fakeDependent))
			Expect(dependents.GetDependent(anotherDependent.ID())).To(Equal(anotherDependent))
		})

		It("should return nil of dependent does not exist", func() {
			Expect(dependents.GetDependent(signaling.NodeID("doesnotexist"))).To(BeNil())
		})
	})

	Context("GetLimit", func() {
		It("should return the limit of the set", func() {
			dependents = signaling.NewDependents(42)
			Expect(dependents.GetLimit()).To(Equal(42))
		})
	})

	Context("GetDependentsCount", func() {
		It("should get dependent count equal to one", func() {
			Expect(dependents.GetDependentsCount()).To(Equal(1))
			Expect(dependents.GetDependent(fakeDependent.ID())).To(Equal(fakeDependent))
		})

		It("should get dependent count equal to two", func() {
			dependents.AddDependent(anotherDependent)

			Expect(dependents.GetDependentsCount()).To(Equal(2))
			Expect(dependents.GetDependent(fakeDependent.ID())).To(Equal(fakeDependent))
		})

		It("should get dependent count equal to zero", func() {
			dependents.RemoveDependent(fakeDependent)
			Expect(dependents.GetDependentsCount()).To(Equal(0))
		})
	})

	Context("AddDependent", func() {
		It("should add dependent", func() {
			Expect(dependents.GetDependent(fakeDependent.ID())).To(Equal(fakeDependent))
			Expect(dependents.GetDependentsCount()).To(Equal(1))
		})

		It("should not add existing dependent", func() {
			dependents.AddDependent(fakeDependent)

			Expect(dependents.GetDependent(fakeDependent.ID())).To(Equal(fakeDependent))
			Expect(dependents.GetDependentsCount()).To(Equal(1))
		})

		It("should inform other dependents of addition", func() {
			dependents.AddDependent(anotherDependent)

			Expect(fakeDependent.ReceiveCallCount()).To(Equal(1))
			message := fakeDependent.ReceiveArgsForCall(0)
			Expect(message.Type).To(Equal(signaling.MessageTypeJoin))
			Expect(message.SourceID).To(Equal(anotherDependent.ID()))
		})
	})

	Context("RemoveDependent", func() {
		It("should remove dependent", func() {
			Expect(dependents.GetDependent(fakeDependent.ID())).To(Equal(fakeDependent))

			dependents.RemoveDependent(fakeDependent)
			Expect(dependents.GetDependent(fakeDependent.ID())).To(BeNil())
			Expect(dependents.GetDependentsCount()).To(Equal(0))
		})

		It("should remove only one dependent", func() {
			dependents.AddDependent(anotherDependent)

			Expect(dependents.GetDependentsCount()).To(Equal(2))

			dependents.RemoveDependent(fakeDependent)
			Expect(dependents.GetDependentsCount()).To(Equal(1))
			Expect(dependents.GetDependent(fakeDependent.ID())).To(BeNil())
			Expect(dependents.GetDependent(anotherDependent.ID())).ToNot(BeNil())
		})

		It("should remove only two dependents", func() {
			anotherDependent = &signalingfakes.FakeReceiverNode{}
			anotherDependent.IDReturns(signaling.NodeID("124"))

			dependents.AddDependent(anotherDependent)

			Expect(dependents.GetDependentsCount()).To(Equal(2))

			dependents.RemoveDependent(fakeDependent)
			Expect(dependents.GetDependentsCount()).To(Equal(1))
			Expect(dependents.GetDependent(fakeDependent.ID())).To(BeNil())
			Expect(dependents.GetDependent(anotherDependent.ID())).ToNot(BeNil())
			Expect(dependents.GetDependent(anotherDependent.ID())).To(Equal(anotherDependent))

			dependents.RemoveDependent(anotherDependent)
			Expect(dependents.GetDependentsCount()).To(Equal(0))
			Expect(dependents.GetDependent(anotherDependent.ID())).To(BeNil())
			Expect(dependents.GetDependent(anotherDependent.ID())).To(BeNil())
		})

		It("should inform other dependents of removal", func() {
			dependents.AddDependent(anotherDependent)

			dependents.RemoveDependent(fakeDependent)

			Expect(anotherDependent.ReceiveCallCount()).To(Equal(1))
			message := anotherDependent.ReceiveArgsForCall(0)
			Expect(message.Type).To(Equal(signaling.MessageTypeLeave))
			Expect(message.SourceID).To(Equal(fakeDependent.ID()))
		})
	})

	Context("MessageDependant", func() {
		It("should message dependant", func() {
			Expect(fakeDependent.ReceiveCallCount()).To(Equal(0))

			dependents.AddDependent(anotherDependent)

			dependents.MessageDependent(signaling.Message{
				Type:          signaling.MessageTypeJoin,
				SourceID:      signaling.NodeID("123"),
				DestinationID: signaling.NodeID("456"),
			})

			// Adding a dependent will call Receive on existing dependents
			Expect(fakeDependent.ReceiveCallCount()).To(Equal(1))
			Expect(anotherDependent.ReceiveCallCount()).To(Equal(1))
		})

		It("should handle trying to message a dependant that does not exist", func() {
			dependents.AddDependent(anotherDependent)

			dependents.MessageDependent(signaling.Message{
				Type:          signaling.MessageTypeJoin,
				SourceID:      signaling.NodeID("123"),
				DestinationID: signaling.NodeID("doesnotexist"),
			})

			// Adding a dependent will call Receive on existing dependents
			Expect(fakeDependent.ReceiveCallCount()).To(Equal(1))
			Expect(anotherDependent.ReceiveCallCount()).To(Equal(0))
		})
	})

	Context("Broadcast", func() {
		It("should message all dependents except source", func() {
			dependents.AddDependent(anotherDependent)

			dependents.Broadcast(signaling.Message{
				Type:     signaling.MessageTypeJoin,
				SourceID: signaling.NodeID("abc"),
			})

			// Adding a dependent will call Receive on existing dependents
			Expect(fakeDependent.ReceiveCallCount()).To(Equal(2))
			Expect(anotherDependent.ReceiveCallCount()).To(Equal(1))
		})
	})
})
