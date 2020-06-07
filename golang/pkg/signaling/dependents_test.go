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
			anotherDependent = &signalingfakes.FakeReceiverNode{}
			anotherDependent.IDReturns(signaling.NodeID("124"))
			dependents.AddDependent(anotherDependent)

			Expect(dependents.GetDependent(fakeDependent.ID())).To(Equal(fakeDependent))
			Expect(dependents.GetDependent(anotherDependent.ID())).To(Equal(anotherDependent))
		})

		It("should get two dependents with unique IDs", func() {
			anotherDependent = &signalingfakes.FakeReceiverNode{}
			anotherDependent.IDReturns(signaling.NodeID("124"))
			anotherDependent.GetParentReturns(fakeParent)
			dependents.AddDependent(anotherDependent)

			Expect(fakeDependent.ID()).ToNot(Equal(anotherDependent.ID()))
		})
	})

	Context("GetDependentsCount", func() {
		It("should get depdenent count equal to one", func() {
			Expect(dependents.GetDependentsCount()).To(Equal(1))
			Expect(dependents.GetDependent(fakeDependent.ID())).To(Equal(fakeDependent))
		})

		It("should get depdenent count equal to two", func() {
			anotherDependent = &signalingfakes.FakeReceiverNode{}
			anotherDependent.IDReturns(signaling.NodeID("124"))
			dependents.AddDependent(anotherDependent)

			Expect(dependents.GetDependentsCount()).To(Equal(2))
			Expect(dependents.GetDependent(fakeDependent.ID())).To(Equal(fakeDependent))
		})

		It("should get depdenent count equal to zero", func() {
			dependents.RemoveDependent(fakeDependent)
			Expect(dependents.GetDependentsCount()).To(Equal(0))
		})
	})

	Context("AddDependent", func() {
		It("should add depdenent", func() {
			Expect(dependents.GetDependent(fakeDependent.ID())).To(Equal(fakeDependent))
			Expect(dependents.GetDependentsCount()).To(Equal(1))
		})

		It("should not add existing depdenent", func() {
			dependents.AddDependent(fakeDependent)

			Expect(dependents.GetDependent(fakeDependent.ID())).To(Equal(fakeDependent))
			Expect(dependents.GetDependentsCount()).To(Equal(1))
		})
	})

	Context("RemoveDependent", func() {
		It("should remove depdenent", func() {
			Expect(dependents.GetDependent(fakeDependent.ID())).To(Equal(fakeDependent))

			dependents.RemoveDependent(fakeDependent)
			Expect(dependents.GetDependent(fakeDependent.ID())).To(BeNil())
			Expect(dependents.GetDependentsCount()).To(Equal(0))
		})

		It("should remove only one depdenent", func() {
			anotherDependent = &signalingfakes.FakeReceiverNode{}
			anotherDependent.IDReturns(signaling.NodeID("124"))

			dependents.AddDependent(fakeDependent)
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
	})

	Context("MessageDepenednt", func() {
		It("should message depdenent", func() {

		})
	})

	Context("Broadcast", func() {
		It("should message all mambers except source", func() {

		})
	})
})
