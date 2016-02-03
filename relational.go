package incata

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb" // MS SQL Server support
	_ "github.com/lib/pq"                // PostgreSQL support
)

// DbType defines the type of the db
type DbType int

// Relational Db Types
const (
	MSSQL      DbType = iota // MS SQL Server
	PostgreSQL               // Postgresql
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
func NewDb(dbType DbType, connection string) (*Db, error) {

	driver, appendStmt, selectStmt, err := getStatements(dbType)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(driver, connection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	database := &Db{
		innerDb:                   db,
		DbType:                    dbType,
		AppendStatement:           appendStmt,
		SelectBySourceIDStatement: selectStmt,
	}
	return database, nil
}

// NewDbFinalized creates a new Db object with setup db argument
func NewDbFinalized(db *sql.DB, dbType DbType) (*Db, error) {

	_, appendStmt, selectStmt, err := getStatements(dbType)
	if err != nil {
		return nil, err
	}

	database := &Db{
		innerDb:                   db,
		DbType:                    dbType,
		AppendStatement:           appendStmt,
		SelectBySourceIDStatement: selectStmt,
	}
	return database, nil
}

func getStatements(dbType DbType) (string, string, string, error) {

	switch dbType {

	case MSSQL:
		return "mssql", `INSERT INTO Event (SourceId, Created, EventType, Version, Payload) VALUES (?, ?, ?, ?, ?)`,
			`SELECT Id ,SourceId ,Created ,EventType ,Version ,Payload FROM Event e WHERE SourceId = ?`, nil

	case PostgreSQL:
		return "postgres", `INSERT INTO linearevents ("SourceId", "Created", "EventType", "Version", "Payload") VALUES ($1, $2, $3, $4, $5)`,
			`SELECT "Id", "SourceId", "Created", "EventType", "Version", "Payload" FROM linearevents WHERE "SourceId" = $1`, nil

	default:
		return "", "", "", fmt.Errorf("DB type %d is not supported", dbType)
	}
}
