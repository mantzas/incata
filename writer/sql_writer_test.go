package writer_test

import (
	"errors"
	"testing"

	"database/sql/driver"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/mantzas/incata/marshal"
	. "github.com/mantzas/incata/mocks"
	. "github.com/mantzas/incata/model"
	. "github.com/mantzas/incata/storage"
	. "github.com/mantzas/incata/writer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"

	"github.com/satori/go.uuid"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

var _ = Describe("Reader", func() {

	It("write data to mocked storage succeeds", func() {

		db, mock, _ := sqlmock.New()
		defer db.Close()

		ser := NewJSONMarshaller()
		storage, _ := NewStorageFinalized(db, MSSQL, "Event")
		wr := NewSQLWriter(storage, ser)

		event := NewEvent(uuid.NewV4(), time.Now(), 1, "TEST", 1)
		payload, _ := ser.Serialize(event.Payload)

		mock.ExpectExec("INSERT INTO Event").WithArgs(event.SourceID.String(), AnyTime{}, "TEST", event.Version, payload).WillReturnResult(sqlmock.NewResult(1, 1))

		err := wr.Write(*event)
		Expect(err).NotTo(HaveOccurred())

		err = mock.ExpectationsWereMet()
		Expect(err).NotTo(HaveOccurred(), "there were unfulfilled expections: %s", err)
	})

	It("write data to mocked storage returns db error", func() {

		db, mock, _ := sqlmock.New()
		defer db.Close()

		ser := NewJSONMarshaller()
		storage, _ := NewStorageFinalized(db, MSSQL, "Event")
		wr := NewSQLWriter(storage, ser)

		event := NewEvent(uuid.NewV4(), time.Now(), 1, "TEST", 1)
		payload, _ := ser.Serialize(event.Payload)

		mock.ExpectExec("INSERT INTO Event").WithArgs(event.SourceID.String(), AnyTime{}, "TEST", event.Version, payload).WillReturnError(errors.New("TEST"))

		err := wr.Write(*event)
		Expect(err).To(HaveOccurred())

		err = mock.ExpectationsWereMet()
		Expect(err).NotTo(HaveOccurred(), "there were unfulfilled expections: %s", err)
	})

	It("write data to mocked storage returns serialization error", func() {

		db, _, _ := sqlmock.New()
		defer db.Close()

		ser := NewJSONMarshaller()
		storage, _ := NewStorageFinalized(db, MSSQL, "Event")
		wr := NewSQLWriter(storage, ser)

		event := NewEvent(uuid.NewV4(), time.Now(), make(map[int]int), "TEST", 1)

		err := wr.Write(*event)
		Expect(err).To(HaveOccurred())
	})

})

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

	event := NewEvent(uuid.NewV4(), time.Now(), GetTestData(), "TEST", 1)

	for n := 0; n < b.N; n++ {
		err := wr.Write(*event)
		if err != nil {
			b.Fatalf("Append error occured! %s", err.Error())
		}
	}
}
