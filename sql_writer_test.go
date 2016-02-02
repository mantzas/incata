package incata

import (
	"testing"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
	"github.com/twinj/uuid"
)

func BenchmarkAppenderPostgresql(b *testing.B) {

	b.Skipf("Postgresql benchmark!")

	db, err := getPostgresqlDb()

	if err != nil {
		b.Fatalf("Fatal error %s", err.Error())
	}

	runDatabaseBenchmark(b, db)
}

func BenchmarkAppenderMsSql(b *testing.B) {

	b.Skipf("SQL Server benchmark!")

	db, err := getSQLServerDb()

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

func getSQLServerDb() (db *Db, err error) {

	db, err = NewDb(MsSQL, "mssql", "Server=xxx;Database=sss;User Id=xx;Password=xxx;",
		`INSERT INTO Event (SourceId, Created, EventType, Version, Payload)  VALUES (?, ?, ?, ?, ?)`,
		`SELECT Id ,SourceId ,Created ,EventType ,Version ,Payload FROM Event e WHERE SourceId = ?`)
	return
}

func getPostgresqlDb() (db *Db, err error) {

	db, err = NewDb(Postgresql, "postgres", "postgres://postgres:xxx@xxx/linear",
		`INSERT INTO linearevents ("SourceId", "Created", "EventType", "Version", "Payload") VALUES ($1, $2, $3, $4, $5)`,
		`SELECT "Id", "SourceId", "Created", "EventType", "Version", "Payload" FROM linearevents WHERE "SourceId" = $1`)
	return
}
