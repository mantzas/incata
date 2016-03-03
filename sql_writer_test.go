package incata

import (
	"errors"
	"testing"

	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/mantzas/incata/mocks"
	. "github.com/mantzas/incata/model"
    . "github.com/mantzas/incata/marshal"

	"github.com/satori/go.uuid"
	"time"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestSqlWriterWrite(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ser := NewJSONMarshaller()
	storage, _ := NewStorageFinalized(db, MSSQL, "Event")
	wr := NewSQLWriter(storage, ser)

	event := NewEvent(uuid.NewV4(), 1, "TEST", 1)
	payload, _ := ser.Serialize(event.Payload)

	mock.ExpectExec("INSERT INTO Event").WithArgs(event.SourceID.String(), AnyTime{}, "TEST", event.Version, payload).WillReturnResult(sqlmock.NewResult(1, 1))

	if err := wr.Write(*event); err != nil {
		t.Errorf("error was not expected while writing event: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestSqlWriterWriteDbError(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ser := NewJSONMarshaller()
	storage, _ := NewStorageFinalized(db, MSSQL, "Event")
	wr := NewSQLWriter(storage, ser)

	event := NewEvent(uuid.NewV4(), 1, "TEST", 1)
	payload, _ := ser.Serialize(event.Payload)

	mock.ExpectExec("INSERT INTO Event").WithArgs(event.SourceID.String(), AnyTime{}, "TEST", event.Version, payload).WillReturnError(errors.New("TEST"))

	err = wr.Write(*event)

	if err == nil {
		t.Fatal("Error was expected!")
	}

	if err.Error() != "TEST" {
		t.Fatalf("Error should have been TEST but was %s", err.Error())
	}
}

func TestSqlWriterWriteSerializationError(t *testing.T) {

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ser := NewJSONMarshaller()
	storage, _ := NewStorageFinalized(db, MSSQL, "Event")
	wr := NewSQLWriter(storage, ser)

	event := NewEvent(uuid.NewV4(), make(map[int]int), "TEST", 1)

	err = wr.Write(*event)

	if err == nil {
		t.Fatalf("error was not expected while writing event: %s", err)
	}

	if err.Error() != "json: unsupported type: map[int]int" {
		t.Fatalf("error was not expected while writing event: %s", err)
	}
}

func BenchmarkAppenderPostgresql(b *testing.B) {

	b.Skipf("Postgresql benchmark!")

	storage, err := NewStorage(PostgreSQL, "postgres://user:pwd@server/linear?sslmode=disable", "linearevents")

	if err != nil {
		b.Fatalf("Fatal error %s", err.Error())
	}

	runDatabaseBenchmark(b, storage)
}

func BenchmarkAppenderMsSql(b *testing.B) {

	b.Skipf("SQL Server benchmark!")

	db, err := NewStorage(MSSQL, "Server=xxx;Database=sss;User Id=xx;Password=xxx;", "Event")

	if err != nil {
		b.Fatalf("Fatal error %s", err.Error())
	}

	runDatabaseBenchmark(b, db)
}

func runDatabaseBenchmark(b *testing.B, storage *Storage) {

	ser := NewJSONMarshaller()
	wr := NewSQLWriter(storage, ser)

	event := NewEvent(uuid.NewV4(), GetTestData(), "TEST", 1)

	for n := 0; n < b.N; n++ {
		err := wr.Write(*event)
		if err != nil {
			b.Fatalf("Append error occured! %s", err.Error())
		}
	}
}
