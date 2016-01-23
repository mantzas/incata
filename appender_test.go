package golinear

import (
	"errors"
	"testing"
	"time"

	"github.com/mantzas/golinear/model"
	"github.com/mantzas/golinear/serializer"
	"github.com/mantzas/golinear/writer"
	"github.com/twinj/uuid"
)

type TestData struct {
	Version   int       `json:"version"`
	Name      string    `json:"name"`
	Balance   float32   `json:"balance"`
	BirthDate time.Time `json:"birth_date"`
}

type TestSerializer struct {
	Failure bool
}

func (serializer TestSerializer) Serialize(value interface{}) (ret string, err error) {

	if serializer.Failure {
		err = errors.New("serialization error")
	} else {
		ret = "Test Value"
	}
	return
}

func TestAppender(t *testing.T) {

	event := model.NewEvent(uuid.NewV4(), getTestData(), "TEST", 1)

	cases := []struct {
		ser         serializer.Serializer
		expectedErr error
	}{
		{
			ser: TestSerializer{
				Failure: false,
			},
			expectedErr: nil,
		},
		{
			ser: TestSerializer{
				Failure: true,
			},
			expectedErr: errors.New("serialization error"),
		},
	}

	for _, c := range cases {

		wr := writer.NewMemoryWriter()
		appender := NewAppender(wr)

		err := appender.Append(*event)

		if err != nil && err.Error() != c.expectedErr.Error() {

			t.Fatalf("Append error occured %s", err)
		} else {

			if len(wr.Data) != 1 {
				t.Fatalf("Expected one item got %d", len(wr.Data))
			}
		}
	}
}

func BenchmarkAppender(b *testing.B) {

	wr := writer.NewMemoryWriter()

	appender := NewAppender(wr)

	event := model.NewEvent(uuid.NewV4(), getTestData(), "TEST", 1)

	for n := 0; n < b.N; n++ {
		appender.Append(*event)
	}
}

func getTestData() *TestData {

	location, _ := time.LoadLocation("Europe/Athens")

	return &TestData{
		Version:   1,
		Name:      "Joe",
		Balance:   12.99,
		BirthDate: time.Date(2015, 12, 13, 23, 59, 59, 0, location),
	}
}
