package incata

import (
	"time"

	. "github.com/mantzas/incata/mocks"
	. "github.com/mantzas/incata/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/satori/go.uuid"
)

var _ = Describe("Appender", func() {

	It("create a new appender without setup", func() {

		SetupAppender(nil)
		appender, err := NewAppender()
		Expect(appender).To(BeNil())
		Expect(err).To(MatchError("Writer is not set up!"))
	})

	It("serialize with error", func() {

		event := NewEvent(uuid.NewV4(), time.Now(), GetTestData(), "TEST", 1)
		var data = make([]Event, 0)
		wr := NewMemoryWriter(data)
		SetupAppender(wr)
		ar, err := NewAppender()
		err = ar.Append(*event)
		Expect(err).NotTo(HaveOccurred())
	})

	Measure("benchmarking appender", func(b Benchmarker) {

		var data = make([]Event, 0)
		wr := NewMemoryWriter(data)
		SetupAppender(wr)
		appender, _ := NewAppender()
		event := NewEvent(uuid.NewV4(), time.Now(), GetTestData(), "TEST", 1)

		runtime := b.Time("runtime", func() {

			appender.Append(*event)
		})

		Expect(runtime.Seconds()).Should(BeNumerically("<", 0.15), "Append() shouldn't take too long.")
	}, 1000)
})
