package writer

import (
	"github.com/mantzas/incata/model"
	"sync"
)

// MemoryWriter Writer for memory
type MemoryWriter struct {
	Data []model.Event
	mx   sync.Mutex
}

// NewMemoryWriter creates a new memory writer
func NewMemoryWriter() *MemoryWriter {

	return &MemoryWriter{
		Data: make([]model.Event, 0),
	}
}

// Write writes a value to a string slice
func (writer *MemoryWriter) Write(event model.Event) (err error) {

	writer.mx.Lock()
	defer writer.mx.Unlock()
	writer.Data = append(writer.Data, event)
	err = nil
	return
}
