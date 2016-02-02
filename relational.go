package incata

import "database/sql"
import "errors"

// DbType defines the type of the db
type DbType int

// Relational Db Types
const (
	MsSQL      DbType = iota // MS SQL Server
	Postgresql               // Postgresql
)

// Db a db struct
type Db struct {
	innerDb                   *sql.DB
	DbType                    DbType
	AppendStatement           string
	SelectBySourceIDStatement string
}

// Exec executes sql statement
func (db *Db) Exec(query string, args ...interface{}) (*sql.Result, error) {
	result, err := db.innerDb.Exec(query, args...)
	return &result, err
}

// Query executes a query statment
func (db *Db) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.innerDb.Query(query, args...)
	return rows, err
}

// Close close db
func (db *Db) Close() (err error) {
	err = db.innerDb.Close()
	return
}

// NewDb return a new MS SQL Server Db object
func NewDb(dbType DbType, driverName string, connection string, appendStatement string, selectBySourceIDStatement string) (database *Db, err error) {

	if len(appendStatement) == 0 {
		err = errors.New("Append statement should have a value!")
		return
	}

	if len(selectBySourceIDStatement) == 0 {
		err = errors.New("Select by source ID statement should have a value!")
		return
	}

	db, err := sql.Open(driverName, connection)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return
	}

	database = &Db{
		innerDb:                   db,
		DbType:                    dbType,
		AppendStatement:           appendStatement,
		SelectBySourceIDStatement: selectBySourceIDStatement,
	}
	return
}
