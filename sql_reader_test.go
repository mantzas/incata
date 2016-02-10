package incata

import (
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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

	rows := sqlmock.NewRows([]string{"Id", "SourceId", "Created", "EventType", "Version", "Payload"}).AddRow(1, uuid.NewV4().String(), time.Now, "Test", 1, "123")

	mock.ExpectQuery("SELECT").WithArgs(sourceID.String()).WillReturnRows(rows)

	storage, _ := NewDbFinalized(db, MSSQL)
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
