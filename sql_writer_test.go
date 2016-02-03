package incata

import (
	"testing"

	"github.com/twinj/uuid"
)

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
