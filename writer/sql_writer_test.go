package writer

import (
	"errors"
	"testing"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
	"github.com/mantzas/incata/model"
	"github.com/mantzas/incata/relational"
	"github.com/mantzas/incata/serializer"
	"github.com/twinj/uuid"
)

type TestData struct {
	Version   int       `json:"version"`
	Name      string    `json:"name"`
	Balance   float32   `json:"balance"`
	BirthDate time.Time `json:"birth_date"`
}

type TestSerializer struct {
	Failure bool
}

func (serializer TestSerializer) Serialize(value interface{}) (ret string, err error) {

	if serializer.Failure {
		err = errors.New("serialization error")
	} else {
		ret = "Test Value"
	}
	return
}

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

func runDatabaseBenchmark(b *testing.B, db *relational.Db) {

	ser := serializer.NewJSONSerializer()
	wr := NewSQLWriter(db, ser)

	event := model.NewEvent(uuid.NewV4(), getTestData(), "TEST", 1)

	for n := 0; n < b.N; n++ {
		err := wr.Write(*event)
		if err != nil {
			b.Fatalf("Append error occured! %s", err.Error())
		}
	}
}

func getSQLServerDb() (db *relational.Db, err error) {

	db, err = relational.NewDb(relational.MsSQL, "mssql", "Server=xxx;Database=sss;User Id=xx;Password=xxx;",
		"INSERT INTO Event (SourceId, Created, EventType, Version, Payload)  VALUES (?, ?, ?, ?, ?)")
	return
}

func getPostgresqlDb() (db *relational.Db, err error) {

	db, err = relational.NewDb(relational.Postgresql, "postgres", "postgres://postgres:xxx@xxx/linear",
		`INSERT INTO linearevents ("SourceId", "Created", "EventType", "Version", "Payload") VALUES ($1, $2, $3, $4, $5)`)

	return
}

func getTestData() *TestData {

	location, _ := time.LoadLocation("Europe/Athens")

	return &TestData{
		Version:   1,
		Name:      "Joe",
		Balance:   12.99,
		BirthDate: time.Date(2015, 12, 13, 23, 59, 59, 0, location),
	}
}
