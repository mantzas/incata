package incata

import (
	"errors"
	"github.com/mantzas/incata/model"
)

// Appender interface
type Appender interface {
	Append(interface{}) error
}

// EventAppender Append events to storage
type EventAppender struct {
	Writer Writer
}

var wr Writer

// SetupAppender setting up the appender
func SetupAppender(writer Writer) {
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
