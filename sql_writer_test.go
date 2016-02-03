package incata

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/twinj/uuid"
)

func TestSqlWriterWrite(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ser := NewJSONSerializer()
	database, _ := NewDbFinalized(db, MSSQL)
	wr := NewSQLWriter(database, ser)

	event := NewEvent(uuid.NewV4(), getTestData(), "TEST", 1)

	payload, _ := ser.Serialize(event.Payload)

	mock.ExpectExec("INSERT INTO Event").WithArgs(event.SourceID, event.Created, event.EventType, event.Version, payload).WillReturnResult(sqlmock.NewResult(1, 1))

	if err := wr.Write(*event); err != nil {
		t.Errorf("error was not expected while writing event: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func BenchmarkAppenderPostgresql(b *testing.B) {

	b.Skipf("Postgresql benchmark!")

	db, err := NewDb(PostgreSQL, "postgres://user:pwd@server/linear?sslmode=disable")

	if err != nil {
		b.Fatalf("Fatal error %s", err.Error())
	}

	runDatabaseBenchmark(b, db)
}

func BenchmarkAppenderMsSql(b *testing.B) {

	b.Skipf("SQL Server benchmark!")

	db, err := NewDb(MSSQL, "Server=xxx;Database=sss;User Id=xx;Password=xxx;")

	if err != nil {
		b.Fatalf("Fatal error %s", err.Error())
	}

	runDatabaseBenchmark(b, db)
}

func runDatabaseBenchmark(b *testing.B, db *Db) {

	ser := NewJSONSerializer()
	wr := NewSQLWriter(db, ser)

	event := NewEvent(uuid.NewV4(), getTestData(), "TEST", 1)

	for n := 0; n < b.N; n++ {
		err := wr.Write(*event)
		if err != nil {
			b.Fatalf("Append error occured! %s", err.Error())
		}
	}
}
