package writer

import (
	"github.com/mantzas/incata/marshal"
	"github.com/mantzas/incata/model"
	"github.com/mantzas/incata/storage"
)

// SQLWriter writer for writing events
type SQLWriter struct {
	Storage    *storage.Storage
	Serializer marshal.Serializer
}

// NewSQLWriter creates a new sql writer
func NewSQLWriter(storage *storage.Storage, ser marshal.Serializer) *SQLWriter {

	return &SQLWriter{Storage: storage, Serializer: ser}
}

// Write writes a value to db table
func (w *SQLWriter) Write(event model.Event) error {

	payloadText, err := w.Serializer.Serialize(event.Payload)
	if err != nil {
		return err
	}

	_, err = w.Storage.Exec(w.Storage.AppendStatement, event.SourceID.String(), event.Created, event.EventType, event.Version, payloadText)
	if err != nil {
		return err
	}

	return nil
}
