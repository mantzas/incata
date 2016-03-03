package incata

import (
	"database/sql/driver"
	"testing"

	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/mantzas/incata/marshal"
	"github.com/satori/go.uuid"
	"time"
)

type AnyType struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyType) Match(v driver.Value) bool {
	return true
}

func TestSqlReaderRead(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var sourceID = uuid.NewV4()

	rows := sqlmock.NewRows([]string{"Id", "SourceId", "Created", "EventType", "Version", "Payload"}).AddRow(1, uuid.NewV4().String(), time.Now(), "Test", 1, "123")
	defer rows.Close()

	mock.ExpectQuery("SELECT").WithArgs(sourceID.String()).WillReturnRows(rows)

	storage, _ := NewStorageFinalized(db, MSSQL, "Event")
	marshaller := NewJSONMarshaller()
	reader := NewSQLReader(storage, marshaller)

	_, err = reader.Read(sourceID)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expections: %s", err)
	}
}

func TestSqlReaderQueryError(t *testing.T) {

	const QUERYERROR = "QUERY ERROR"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var sourceID = uuid.NewV4()

	mock.ExpectQuery("SELECT").WithArgs(sourceID.String()).WillReturnError(errors.New(QUERYERROR))

	storage, _ := NewStorageFinalized(db, MSSQL, "Event")
	marshaller := NewJSONMarshaller()
	reader := NewSQLReader(storage, marshaller)

	_, err = reader.Read(sourceID)
	if err == nil {
		t.Fatalf("unexpected err: %s", err)
	}

	if err.Error() != QUERYERROR {
		t.Fatalf("there were unfulfilled expections: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expections: %s", err)
	}
}
