package incata

import (
	"github.com/satori/go.uuid"
)

// SQLReader for reading events
type SQLReader struct {
	Storage    *Storage
	Serializer Serializer
}

// NewSQLReader creates a new sql reader
func NewSQLReader(storage *Storage, ser Serializer) *SQLReader {

	return &SQLReader{Storage: storage, Serializer: ser}
}

// Read from db with source id
func (r *SQLReader) Read(sourceID uuid.UUID) ([]Event, error) {

	rows, err := r.Storage.Query(r.Storage.SelectBySourceIDStatement, sourceID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events = make([]Event, 0)

	for rows.Next() {
		var event = new(Event)

		if err := rows.Scan(&event.ID, &event.SourceID, &event.Created, &event.EventType, &event.Version, &event.Payload); err != nil {
			return nil, err
		}

		events = append(events, *event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
