package incata

import (
	"errors"
	"github.com/mantzas/incata/model"
	"github.com/mantzas/incata/reader"
	"github.com/satori/go.uuid"
)

// Retriever interface
type Retriever interface {
	Retrieve(uuid.UUID) ([]model.Event, error)
}

// EventRetriever Append events to storage
type EventRetriever struct {
	Reader reader.Reader
}

var rd reader.Reader

// SetupRetriever setting up the retriever
func SetupRetriever(reader reader.Reader) {
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
func (appender *EventRetriever) Retrieve(sourceID uuid.UUID) ([]model.Event, error) {

	return appender.Reader.Read(sourceID)
}
