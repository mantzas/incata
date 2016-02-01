package writer

import (
	"github.com/mantzas/incata/model"
	"github.com/mantzas/incata/relational"
	"github.com/mantzas/incata/serializer"
)

// SQLWriter writer for MS SQL Server
type SQLWriter struct {
	Db         *relational.Db
	Serializer serializer.Serializer
}

// NewSQLWriter creates a new sql writer
func NewSQLWriter(db *relational.Db, ser serializer.Serializer) *SQLWriter {

	return &SQLWriter{Db: db, Serializer: ser}
}

// Write writes a value to db table
func (writer *SQLWriter) Write(event model.Event) error {

	payloadText, err := writer.Serializer.Serialize(event.Payload)
	if err != nil {
		return err
	}

	_, err = writer.Db.Exec(writer.Db.AppendStatement, event.SourceID.String(), event.Created, event.EventType, event.Version, payloadText)
	if err != nil {
		return err
	}

	return nil
}
