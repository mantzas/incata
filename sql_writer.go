package incata

// SQLWriter writer for writing events
type SQLWriter struct {
	Db         *Db
	Serializer Serializer
}

// NewSQLWriter creates a new sql writer
func NewSQLWriter(db *Db, ser Serializer) *SQLWriter {

	return &SQLWriter{Db: db, Serializer: ser}
}

// Write writes a value to db table
func (w *SQLWriter) Write(event Event) error {

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
