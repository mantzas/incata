package writer

import "github.com/mantzas/incata/model"

// Writer Interface for writing events to storage
type Writer interface {
	Write(model.Event) error
}
