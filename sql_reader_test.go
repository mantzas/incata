package incata

import (
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/twinj/uuid"
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

	storage, _ := NewDbFinalized(db, MSSQL)

	var sourceID = uuid.NewV4()

	mock.ExpectQuery("SELECT Id ,SourceId ,Created ,EventType ,Version ,Payload FROM Event WHERE SourceId =").WithArgs(AnyTime{})

	marshaller := NewJSONMarshaller()
	reader := NewSQLReader(storage, marshaller)

	reader.Read(sourceID)

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expections: %s", err)
	}
}
