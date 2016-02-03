package incata

import (
	"errors"
	"github.com/twinj/uuid"
)

// Retriever interface
type Retriever interface {
	Retrieve(uuid.UUID) ([]Event, error)
}

// EventRetriever Append events to storage
type EventRetriever struct {
	Reader Reader
}

var rd Reader

// SetupRetriever setting up the retriever
func SetupRetriever(reader Reader) {
	rd = reader
}

// NewRetriever creates a new event retriever
func NewRetriever() (*EventRetriever, error) {

	if rd == nil {
		return nil, errors.New("Reader is not set up!")
	}
	return &EventRetriever{Reader: rd}, nil
}

// Retrieve  events based on Source ID
func (appender *EventRetriever) Retrieve(sourceID uuid.UUID) ([]Event, error) {

	return appender.Reader.Read(sourceID)
}
