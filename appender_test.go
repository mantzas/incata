package incata

import (
	"testing"

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

		event := NewEvent(uuid.NewV4(), GetTestData(), "TEST", 1)
		var data = make([]Event, 0)
		wr := NewMemoryWriter(data)
		SetupAppender(wr)
		ar, err := NewAppender()
		err = ar.Append(*event)
		Expect(err).NotTo(HaveOccurred())
	})

})

func BenchmarkAppender(b *testing.B) {

	var data = make([]Event, 0)
	wr := NewMemoryWriter(data)

	SetupAppender(wr)

	appender, _ := NewAppender()

	event := NewEvent(uuid.NewV4(), GetTestData(), "TEST", 1)

	for n := 0; n < b.N; n++ {
		appender.Append(*event)
	}
}
