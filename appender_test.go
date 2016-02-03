package incata

import (
	"errors"
	"testing"
	"time"

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

func (serializer TestSerializer) Serialize(value interface{}) (ret interface{}, err error) {

	if serializer.Failure {
		err = errors.New("serialization error")
	} else {
		ret = "Test Value"
	}
	return
}

func TestNewAppenderWithoutSetup(t *testing.T) {

	_, err := NewAppender()
	if err == nil {
		t.Fatal(err.Error())
	}
}

func TestAppender(t *testing.T) {

	event := NewEvent(uuid.NewV4(), getTestData(), "TEST", 1)

	cases := []struct {
		ser         Serializer
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

		var data = make([]Event, 0)
		wr := NewMemoryWriter(data)
		SetupAppender(wr)
		ar, err := NewAppender()
		if err != nil {
			t.Fatal(err.Error())
		}

		err = ar.Append(*event)

		if err != nil && err.Error() != c.expectedErr.Error() {

		} else {

			if len(wr.Data) != 1 {
				t.Fatalf("Expected one item got %d", len(wr.Data))
			}
		}
	}
}

func BenchmarkAppender(b *testing.B) {

	var data = make([]Event, 0)
	wr := NewMemoryWriter(data)

	SetupAppender(wr)

	appender, _ := NewAppender()

	event := NewEvent(uuid.NewV4(), getTestData(), "TEST", 1)

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
