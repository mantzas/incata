package incata

import (
	"github.com/satori/go.uuid"
	"testing"
)

func TestRetrieverWithoutSetup(t *testing.T) {
	_, err := NewRetriever()

	if err == nil {
		t.Fatal("Error should have occured!")
	}
}

func TestRetriever(t *testing.T) {

	var sourceID = uuid.NewV4()
	var data = make([]Event, 0)

	data = append(data, *NewEvent(uuid.NewV4(), getTestData(), "TEST", 1))
	data = append(data, *NewEvent(sourceID, getTestData(), "TEST", 1))
	data = append(data, *NewEvent(uuid.NewV4(), getTestData(), "TEST", 1))
	data = append(data, *NewEvent(sourceID, getTestData(), "TEST", 1))
	data = append(data, *NewEvent(uuid.NewV4(), getTestData(), "TEST", 1))

	rd := NewMemoryReader(data)

	SetupRetriever(rd)

	r, err := NewRetriever()

	if err != nil {
		t.Fatal("Error getting new retriever!")
	}

	events, err := r.Retrieve(sourceID)

	if err != nil {
		t.Fatal("Error retrieving events!")
	}

	if len(events) != 2 {
		t.Fatalf("Expected 2 events but was %d", len(events))
	}
}

// MemoryReader for memory reading
type MemoryReader struct {
	Data []Event
}

// NewMemoryReader creates a new memory reader
func NewMemoryReader(data []Event) *MemoryReader {

	return &MemoryReader{
		Data: data,
	}
}

// Write writes a value to a string slice
func (r *MemoryReader) Read(sourceID uuid.UUID) ([]Event, error) {

	var events = make([]Event, 0)

	for _, event := range r.Data {

		if event.SourceID == sourceID {
			events = append(events, event)
		}
	}

	return events, nil
}
