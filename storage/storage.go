package storage

import (
	"database/sql"
	"fmt"
	"strings"
)

// DbType defines the type of the db
type DbType int

// Relational Db Types
const (
	MSSQL      DbType = iota // MS SQL Server
	PostgreSQL               // Postgresql
)

var dbTypeMap = map[string]DbType{
	"mssql":      MSSQL,
	"postgresql": PostgreSQL,
}

// ConvertToDbType convert's a string to a DbType
func ConvertToDbType(value string) (DbType, error) {

	dbType, ok := dbTypeMap[strings.ToLower(value)]

	if ok {
		return dbType, nil
	}

	return 0, fmt.Errorf("Failed to convert %s to db type", value)
}

// Storage a db abstraction
type Storage struct {
	innerDb                   *sql.DB
	DbType                    DbType
	AppendStatement           string
	SelectBySourceIDStatement string
}

// Exec executes sql statement
func (db *Storage) Exec(query string, args ...interface{}) (*sql.Result, error) {
	result, err := db.innerDb.Exec(query, args...)
	return &result, err
}

// Query executes a query statement
func (db *Storage) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.innerDb.Query(query, args...)
	return rows, err
}

// Close close db
func (db *Storage) Close() (err error) {
	err = db.innerDb.Close()
	return
}

// NewStorage creates a new storage
func NewStorage(dbType DbType, connection string, tableName string) (*Storage, error) {

	driver, appendStmt, selectStmt, err := getStatements(dbType, tableName)
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

	storage := &Storage{
		innerDb:                   db,
		DbType:                    dbType,
		AppendStatement:           appendStmt,
		SelectBySourceIDStatement: selectStmt,
	}
	return storage, nil
}

// NewStorageFinalized creates a new storage with a passed in db argument
func NewStorageFinalized(db *sql.DB, dbType DbType, tableName string) (*Storage, error) {

	_, appendStmt, selectStmt, err := getStatements(dbType, tableName)
	if err != nil {
		return nil, err
	}

	storage := &Storage{
		innerDb:                   db,
		DbType:                    dbType,
		AppendStatement:           appendStmt,
		SelectBySourceIDStatement: selectStmt,
	}
	return storage, nil
}

func getStatements(dbType DbType, tableName string) (string, string, string, error) {

	switch dbType {

	case MSSQL:
		return "mssql", fmt.Sprintf("INSERT INTO %s (SourceId, Created, EventType, Version, Payload) VALUES (?, ?, ?, ?, ?)", tableName),
			fmt.Sprintf("SELECT Id, CAST(SourceId AS CHAR(36)), Created, EventType, Version, Payload FROM %s WHERE SourceId = ?", tableName), nil

	case PostgreSQL:
		return "postgres", fmt.Sprintf(`insert into %s ("source_id","created","event_type","version","payload") values ($1, $2, $3, $4, $5)`, tableName),
			fmt.Sprintf(`SELECT id, source_id, created, event_type, version, payload FROM %s WHERE source_id = $1`, tableName), nil

	default:
		return "", "", "", fmt.Errorf("DB type %d is not supported", dbType)
	}
}
