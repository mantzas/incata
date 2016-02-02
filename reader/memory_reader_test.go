package reader

import (
	"github.com/mantzas/incata/model"
	"github.com/twinj/uuid"
	"testing"
)

func TestReadItems(t *testing.T) {

	var data = make([]model.Event, 0)
	sourceID := uuid.NewV4()
	data = append(data, *getEvent(sourceID))
	data = append(data, *getEvent(uuid.NewV4()))
	data = append(data, *getEvent(sourceID))
	data = append(data, *getEvent(uuid.NewV4()))
	data = append(data, *getEvent(sourceID))
	data = append(data, *getEvent(uuid.NewV4()))

	reader := NewMemoryReader(data)

	events, _ := reader.Read(sourceID)

	if len(events) != 3 {
		t.Fatalf("Expected 3 actual %d", len(events))
	}

	for _, event := range events {
		if event.SourceID != sourceID {
			t.Fatalf("SourceID was not expected!")
		}
	}
}

func getEvent(sourceID uuid.UUID) *model.Event {

	var event = new(model.Event)

	event.SourceID = sourceID

	return event
}
