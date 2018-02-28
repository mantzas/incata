package reader_test

import (
	"database/sql/driver"

	"errors"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/mantzas/incata/marshal"
	. "github.com/mantzas/incata/reader"
	. "github.com/mantzas/incata/storage"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/satori/go.uuid"
)

type AnyType struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyType) Match(v driver.Value) bool {
	return true
}

var _ = Describe("Reader", func() {

	It("read data from mocked storage succeeds", func() {

		db, mock, _ := sqlmock.New()

		defer db.Close()

		var sourceID = uuid.NewV4()

		rows := sqlmock.NewRows([]string{"Id", "SourceId", "Created", "EventType", "Version", "Payload"}).AddRow(1, uuid.NewV4().String(), time.Now(), "Test", 1, "123")

		mock.ExpectQuery("SELECT").WithArgs(sourceID.String()).WillReturnRows(rows)

		storage, _ := NewStorageFinalized(db, MSSQL, "Event")
		marshaller := NewJSONMarshaller()
		reader := NewSQLReader(storage, marshaller)

		_, err := reader.Read(sourceID)
		Expect(err).NotTo(HaveOccurred())

		err = mock.ExpectationsWereMet()
		Expect(err).NotTo(HaveOccurred(), "there were unfulfilled expections: %s", err)
	})

	It("read data from mocked storage returns query error", func() {

		const QUERYERROR = "QUERY ERROR"

		db, mock, _ := sqlmock.New()
		defer db.Close()

		var sourceID = uuid.NewV4()

		mock.ExpectQuery("SELECT").WithArgs(sourceID.String()).WillReturnError(errors.New(QUERYERROR))

		storage, _ := NewStorageFinalized(db, MSSQL, "Event")
		marshaller := NewJSONMarshaller()
		reader := NewSQLReader(storage, marshaller)

		_, err := reader.Read(sourceID)
		Expect(err).To(MatchError(QUERYERROR))

		err = mock.ExpectationsWereMet()
		Expect(err).NotTo(HaveOccurred(), "there were unfulfilled expections: %s", err)
	})
})
