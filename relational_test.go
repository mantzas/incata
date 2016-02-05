package incata

import (
	"testing"

	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestDbExec(t *testing.T) {

	dbInner, mock, _ := sqlmock.New()

	db, _ := NewDbFinalized(dbInner, MSSQL)

	mock.ExpectExec("123").WithArgs(1, 2, 3).WillReturnResult(sqlmock.NewResult(1, 1))

	db.Exec("123", 1, 2, 3)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expectations not met, %s", err)
	}
}

func TestDbQuery(t *testing.T) {

	dbInner, mock, _ := sqlmock.New()

	db, _ := NewDbFinalized(dbInner, PostgreSQL)

	var rows driver.Rows

	mock.ExpectQuery("123").WithArgs(1, 2, 3).WillReturnRows(rows)

	db.Query("123", 1, 2, 3)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expectations not met, %s", err)
	}
}

func TestDbClose(t *testing.T) {

	dbInner, mock, _ := sqlmock.New()

	db, _ := NewDbFinalized(dbInner, MSSQL)

	mock.ExpectClose()

	db.Close()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expectations not met, %s", err)
	}
}
