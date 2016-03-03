package incata

import "github.com/mantzas/incata/model"

// SQLWriter writer for writing events
type SQLWriter struct {
	Storage    *Storage
	Serializer Serializer
}

// NewSQLWriter creates a new sql writer
func NewSQLWriter(storage *Storage, ser Serializer) *SQLWriter {

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
