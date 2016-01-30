package relational

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
	innerDb         *sql.DB
	DbType          DbType
	AppendStatement string
}

// Exec executes sql statement
func (db *Db) Exec(query string, args ...interface{}) (*sql.Result, error) {
	result, err := db.innerDb.Exec(query, args...)
	return &result, err
}

// Close close db
func (db *Db) Close() (err error) {
	err = db.innerDb.Close()
	return
}

// NewDb return a new MS SQL Server Db object
func NewDb(dbType DbType, driverName string, connection string, appendStatement string) (database *Db, err error) {

	if len(appendStatement) == 0 {
		err = errors.New("Append statement should have a value!")
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
		innerDb:         db,
		DbType:          dbType,
		AppendStatement: appendStatement,
	}
	return
}
