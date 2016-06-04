package incata

import (
	"errors"

	"github.com/mantzas/incata/model"
	"github.com/mantzas/incata/writer"
)

// Appender interface
type Appender interface {
	Append(event model.Event) error
}

// EventAppender Append events to storage
type EventAppender struct {
	Writer writer.Writer
}

var wr writer.Writer

// SetupAppender setting up the appender
func SetupAppender(writer writer.Writer) {
	wr = writer
}

// NewAppender Creates a new event appender
func NewAppender() (*EventAppender, error) {

	if wr == nil {
		return nil, errors.New("Writer is not set up!")
	}
	return &EventAppender{Writer: wr}, nil
}

// Append Append the payload to the storage
func (appender *EventAppender) Append(event model.Event) error {

	return appender.Writer.Write(event)
}
