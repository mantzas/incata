package marshal_test

import (
	"testing"
	"time"

	. "github.com/mantzas/incata/marshal"
	. "github.com/mantzas/incata/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Marshal", func() {

	It("serialize test data to json and match", func() {
		expectedString := `{"version":1,"name":"Joe","balance":12.99,"birth_date":"2015-12-13T23:59:59+02:00"}`

		location, _ := time.LoadLocation("Europe/Athens")

		testData := TestData{
			Version:   1,
			Name:      "Joe",
			Balance:   12.99,
			BirthDate: time.Date(2015, 12, 13, 23, 59, 59, 0, location),
		}

		serializedString, err := NewJSONMarshaller().Serialize(testData)

		Expect(serializedString).To(Equal(expectedString))
		Expect(err).NotTo(HaveOccurred())
	})

	It("serialize unsupported data type fails", func() {
		var m = make(map[int]int, 0)
		_, err := NewJSONMarshaller().Serialize(m)

		Expect(err).To(HaveOccurred())
	})

	It("deserialize json to test data and match", func() {
		location, _ := time.LoadLocation("Europe/Athens")

		expected := TestData{
			Version:   1,
			Name:      "Joe",
			Balance:   12.99,
			BirthDate: time.Date(2015, 12, 13, 23, 59, 59, 0, location),
		}

		actualData := `{"version":1,"name":"Joe","balance":12.99,"birth_date":"2015-12-13T23:59:59+02:00"}`
		var actual TestData

		err := NewJSONMarshaller().Deserialize(actualData, &actual)

		Expect(actual.Balance).To(Equal(expected.Balance))
		Expect(actual.BirthDate.Equal(expected.BirthDate)).To(BeTrue())
		Expect(actual.Name).To(Equal(expected.Name))
		Expect(actual.Version).To(Equal(expected.Version))
		Expect(err).NotTo(HaveOccurred())
	})

	It("deserialize fails due to invalid json", func() {

		var actual TestData
		err := NewJSONMarshaller().Deserialize(`{"version":1,"name":"Joe","balance":12.99,"birth_date":"2015-12-13T23:59:59+02:00------"}`, &actual)
		Expect(err).To(HaveOccurred())
	})

	It("deserialize wrong to the struct", func() {

		var actual TestData
		err := NewJSONMarshaller().Deserialize(123, &actual)
		Expect(err).To(HaveOccurred())
	})
})

func BenchmarkJSONSerializer(b *testing.B) {

	var sert = NewJSONMarshaller()

	location, err := time.LoadLocation("Europe/Athens")

	if err != nil {
		b.Fatalf("Error getting location!")
	}

	testData := TestData{
		Version:   1,
		Name:      "Joe",
		Balance:   12.99,
		BirthDate: time.Date(2015, 12, 13, 23, 59, 59, 0, location),
	}

	for n := 0; n < b.N; n++ {
		sert.Serialize(testData)
	}
}
