package reader

import (
	"github.com/mantzas/incata/model"
	"github.com/mantzas/incata/relational"
	"github.com/mantzas/incata/serializer"
	"github.com/twinj/uuid"
)

// SQLReader for reading events
type SQLReader struct {
	Db         *relational.Db
	Serializer serializer.Serializer
}

// NewSQLReader creates a new sql reader
func NewSQLReader(db *relational.Db, ser serializer.Serializer) *SQLReader {

	return &SQLReader{Db: db, Serializer: ser}
}

// Read from db with source id
func (r *SQLReader) Read(sourceID uuid.UUID) ([]model.Event, error) {

	rows, err := r.Db.Query(r.Db.SelectBySourceIDStatement, sourceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events = make([]model.Event, 0)

	for rows.Next() {
		var event = new(model.Event)

		if err := rows.Scan(&event.Id, &event.SourceID, &event.Created, &event.Payload, &event.EventType, &event.Version); err != nil {
			return nil, err
		}
		events = append(events, *event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
