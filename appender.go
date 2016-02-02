package incata

import (
	"errors"
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
		return nil, errors.New("Writer is not setup!")
	}
	return &EventAppender{Writer: wr}, nil
}

// Append Append the payload to the storage
func (appender *EventAppender) Append(event Event) error {

	return appender.Writer.Write(event)
}
