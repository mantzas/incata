package reader

import (
	"github.com/mantzas/incata/model"
	"github.com/twinj/uuid"
)

// Reader interface for getting events
type Reader interface {
	Read(uuid.UUID) ([]model.Event, error)
}
