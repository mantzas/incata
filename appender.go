package golinear

import (
	"github.com/mantzas/golinear/model"
	"github.com/mantzas/golinear/writer"
)

// Appender interface
type Appender interface {
	Append(interface{}) error
}

// EventAppender Append events to storage
type EventAppender struct {
	Writer writer.Writer
}

// NewAppender Creates a new event appender
func NewAppender(writer writer.Writer) *EventAppender {
	return &EventAppender{
		Writer: writer,
	}
}

// Append Append the payload to the storage
func (appender *EventAppender) Append(event model.Event) error {

	return appender.Writer.Write(event)
}
