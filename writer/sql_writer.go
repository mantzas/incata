package writer

import (
	"github.com/mantzas/incata/model"
	"github.com/mantzas/incata/relational"
	"github.com/mantzas/incata/serializer"
)

// SQLWriter writer for writing events
type SQLWriter struct {
	Db         *relational.Db
	Serializer serializer.Serializer
}

// NewSQLWriter creates a new sql writer
func NewSQLWriter(db *relational.Db, ser serializer.Serializer) *SQLWriter {

	return &SQLWriter{Db: db, Serializer: ser}
}

// Write writes a value to db table
func (w *SQLWriter) Write(event model.Event) error {

	payloadText, err := w.Serializer.Serialize(event.Payload)
	if err != nil {
		return err
	}

	_, err = w.Db.Exec(w.Db.AppendStatement, event.SourceID.String(), event.Created, event.EventType, event.Version, payloadText)
	if err != nil {
		return err
	}

	return nil
}
