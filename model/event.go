package model

import (
	"github.com/twinj/uuid"
	"time"
)

// Event this is the main event that will get written
type Event struct {
	SourceID  uuid.UUID
	Created   time.Time
	Payload   interface{}
	EventType string
	Version   int
}

// NewEvent creating a new event
func NewEvent(sourceID uuid.UUID, payload interface{}, eventType string, version int) *Event {
	return &Event{
		SourceID:  sourceID,
		Created:   time.Now().UTC(),
		Payload:   payload,
		EventType: eventType,
		Version:   version,
	}
}
