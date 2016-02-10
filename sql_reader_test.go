package incata

import (
	"database/sql/driver"
	"testing"

	//"github.com/DATA-DOG/go-sqlmock"
	"github.com/satori/go.uuid"
)

type AnyType struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyType) Match(v driver.Value) bool {
	return true
}

func TestSqlReaderRead(t *testing.T) {

	t.SkipNow()
	//
	// 	db, mock, err := sqlmock.New()
	// 	if err != nil {
	// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	// 	}
	// 	defer db.Close()
	//
	// 	var sourceID = uuid.NewV4()
	//
	// 	rows := sqlmock.NewRows([]string{"Id", "SourceId", "Created", "EventType", "Version", "Payload"})
	//
	// 	mock.ExpectQuery("SELECT").WithArgs(sourceID.String()).WillReturnRows(rows)
	//
	// 	storage, _ := NewDbFinalized(db, MSSQL)
	// 	marshaller := NewJSONMarshaller()
	// 	reader := NewSQLReader(storage, marshaller)

	// 	_, err = reader.Read(sourceID)
	// 	if err != nil {
	// 		t.Fatalf("unexpected err: %s", err)
	// 	}
	//
	// 	if err = mock.ExpectationsWereMet(); err != nil {
	// 		t.Fatalf("there were unfulfilled expections: %s", err)
	// 	}
}

func TestUnmarshalUUID(t *testing.T) {

	srcID := uuid.NewV4()

	srcIDBytes := srcID.Bytes()

	var srcIDResult uuid.UUID

	srcIDResult, _ = uuid.FromBytes(srcIDBytes)

	if !uuid.Equal(srcID, srcIDResult) {
		t.Fatalf("are not equal %s %s", srcID, srcIDResult)
	}
}
