package incata

import (
	"github.com/twinj/uuid"
)

// Reader interface for getting events
type Reader interface {
	Read(uuid.UUID) ([]Event, error)
}
