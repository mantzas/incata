package incata

import (
	"github.com/twinj/uuid"
)

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
