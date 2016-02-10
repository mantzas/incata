package incata

import (
	"testing"

	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestStorageExec(t *testing.T) {

	dbInner, mock, _ := sqlmock.New()

	storage, _ := NewStorageFinalized(dbInner, MSSQL, "Event")

	mock.ExpectExec("123").WithArgs(1, 2, 3).WillReturnResult(sqlmock.NewResult(1, 1))

	storage.Exec("123", 1, 2, 3)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expectations not met, %s", err)
	}
}

func TestStorageQuery(t *testing.T) {

	dbInner, mock, _ := sqlmock.New()

	storage, _ := NewStorageFinalized(dbInner, PostgreSQL, "Event")

	var rows driver.Rows

	mock.ExpectQuery("123").WithArgs(1, 2, 3).WillReturnRows(rows)

	storage.Query("123", 1, 2, 3)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expectations not met, %s", err)
	}
}

func TestStorageClose(t *testing.T) {

	dbInner, mock, _ := sqlmock.New()

	storage, _ := NewStorageFinalized(dbInner, MSSQL, "Event")

	mock.ExpectClose()

	storage.Close()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Expectations not met, %s", err)
	}
}
