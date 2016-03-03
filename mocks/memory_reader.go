package mocks

import (
	"github.com/mantzas/incata/model"
	"github.com/satori/go.uuid"
)

// MemoryReader for memory reading
type MemoryReader struct {
	Data []model.Event
}

// NewMemoryReader creates a new memory reader
func NewMemoryReader(data []model.Event) *MemoryReader {

	return &MemoryReader{
		Data: data,
	}
}

// Write writes a value to a string slice
func (r *MemoryReader) Read(sourceID uuid.UUID) ([]model.Event, error) {

	var events = make([]model.Event, 0)

	for _, event := range r.Data {

		if event.SourceID == sourceID {
			events = append(events, event)
		}
	}

	return events, nil
}
