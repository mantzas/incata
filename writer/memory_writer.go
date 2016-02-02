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
func NewMemoryWriter(data []model.Event) *MemoryWriter {

	return &MemoryWriter{
		Data: data,
	}
}

// Write writes a value to a string slice
func (w *MemoryWriter) Write(event model.Event) (err error) {

	w.mx.Lock()
	defer w.mx.Unlock()
	w.Data = append(w.Data, event)
	err = nil
	return
}
