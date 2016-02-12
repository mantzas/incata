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

func TestStorageWrongDbType(t *testing.T) {

	dbInner, _, _ := sqlmock.New()

	_, err := NewStorageFinalized(dbInner, 3, "Event")

	if err == nil {
        t.Fatal("Should have been returning a error")
    }
}

func TestNewStorageWrongDbType(t *testing.T) {

	_, err := NewStorage(3, "123" ,"Event")

	if err == nil {
        t.Fatal("Should have been returning a error")
    }
}

func TestNewStorageOpenError(t *testing.T) {

	_, err := NewStorage(MSSQL, "123" ,"Event")

	if err == nil {
        t.Fatal("Should have been returning a error")
    }
}

var convertToTypeTests = []struct {
    in  string
    out DbType
    hasError bool
}{
    {"mssql", MSSQL, false},
    {"MSSQL", MSSQL, false},
    {"MsSQL", MSSQL, false},
    {"MsSql", MSSQL, false},
    {"postgresql", PostgreSQL, false},
    {"PostgreSQL", PostgreSQL, false},
    {"POSTGRESQL", PostgreSQL, false},
    {"xxx", PostgreSQL, true},
}

func TestConvertToType(t *testing.T)  {
    
    for _, test := range convertToTypeTests {
        
        actual, err := ConvertToDbType(test.in)
        
        if err != nil && test.hasError {
            
            if(!test.hasError){
                t.Errorf("Expected no errors but shit happens for input %s", test.in)    
            }            
        } else if actual != test.out {
            t.Errorf("Expected %d got %d for input %s", test.out, actual, test.in)
        }
    }
}