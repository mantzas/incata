package incata

import (
	"github.com/twinj/uuid"
	"testing"
	"time"
)

func TestNewEvent(t *testing.T) {

	username := "user name"
	event := NewEvent(uuid.NewV4(), username, "test type", 1)

	if event.Version != 1 {

		t.Fatalf("Version number wrong! %d", event.Version)
	}

	if event.EventType != "test type" {

		t.Fatalf("EventType number wrong! %s", event.EventType)
	}

	if event.Payload != username {

		t.Fatalf("Payload wrong! %s", event.Payload)
	}

	utcNow := time.Now().UTC()

	if event.Created.After(utcNow) {

		t.Fatalf("Time is not less or equal! %s", event.Created)
	}

	if event.SourceID.Version() != int(uuid.RFC4122v4) {

		t.Fatalf("uuid version is not 4! %d", event.SourceID.Version())
	}
}
