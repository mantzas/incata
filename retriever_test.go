package incata

import (
	. "github.com/mantzas/incata/mocks"
	. "github.com/mantzas/incata/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/satori/go.uuid"
)

var _ = Describe("Retriever", func() {

	It("create a new appender without setup", func() {

		SetupRetriever(nil)
		retriever, err := NewRetriever()
		Expect(retriever).To(BeNil())
		Expect(err).To(MatchError("Reader is not set up!"))
	})

	It("retrieve data succeeds", func() {

		var sourceID = uuid.NewV4()
		var data = make([]Event, 0)

		data = append(data, *NewEvent(uuid.NewV4(), GetTestData(), "TEST", 1))
		data = append(data, *NewEvent(sourceID, GetTestData(), "TEST", 1))
		data = append(data, *NewEvent(uuid.NewV4(), GetTestData(), "TEST", 1))
		data = append(data, *NewEvent(sourceID, GetTestData(), "TEST", 1))
		data = append(data, *NewEvent(uuid.NewV4(), GetTestData(), "TEST", 1))

		rd := NewMemoryReader(data)

		SetupRetriever(rd)

		r, err := NewRetriever()
		Expect(err).NotTo(HaveOccurred())

		events, err := r.Retrieve(sourceID)
		Expect(err).NotTo(HaveOccurred())

		Expect(len(events)).To(Equal(2))
	})
})
