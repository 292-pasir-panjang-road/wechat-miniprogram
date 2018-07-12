package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	POSTGRES         = "postgres"
	DOES_NOT_SUPPORT = "database: PsqlDB cannot initialize other types of DB."

	CONNECTION_DEFAULT_TEMPLATE = "host=%s port=%s dbname=%s user=%s password='%s'"
	CONNECTION_GAE_TEMPLATE     = "host=%s dbname=%s user=%s password='%s'"
	CONNECTION_NO_SSL           = " sslmode=disable"
)

// Create a new db struct
type PsqlDB struct {
	connection *sqlx.DB
}

// Creates a new psql database (most basic and traditional way)
func New(config DBConfig) (Database, error) {

	// Only allows postgres
	if config.Driver != POSTGRES {
		return nil, errors.New(DOES_NOT_SUPPORT)
	}

	// Gets connection string
	connectionString := fmt.Sprintf(CONNECTION_DEFAULT_TEMPLATE,
		config.Host, config.Port, config.DBName, config.Username, config.Password,
	)

	if !config.EnableSSL {
		connectionString += CONNECTION_NO_SSL
	}

	// Makes connection
	connection, err := sqlx.Connect(POSTGRES, connectionString)
	if err != nil {
		return nil, err
	}

	connection.SetMaxIdleConns(config.MaxIdleConnections)
	connection.SetMaxOpenConns(config.MaxOpenConnections)

	return &PsqlDB{connection}, nil
}

// Creates a new psql database (for Google App Engine only)
func NewGAE(config DBConfig) (Database, error) {

	// Only allows postgres
	if config.Driver != POSTGRES {
		return nil, errors.New(DOES_NOT_SUPPORT)
	}

	connectionString := fmt.Sprintf(CONNECTION_GAE_TEMPLATE,
		config.Host, config.DBName, config.Username, config.Password,
	)

	if !config.EnableSSL {
		connectionString += CONNECTION_NO_SSL
	}

	// Makes connection
	connection, err := sqlx.Open(POSTGRES, connectionString)
	if err != nil {
		return nil, err
	}

	connection.SetMaxIdleConns(config.MaxIdleConnections)
	connection.SetMaxOpenConns(config.MaxOpenConnections)

	return &PsqlDB{connection}, nil
}

// Selects one item. result in "out" (casted)
func (d *PsqlDB) SelectOne(out interface{}, sql string, args ...interface{}) error {
	return d.connection.Get(out, sql, args...)
}

// Selects several items. result in "out" (casted)
func (d *PsqlDB) SelectMany(out interface{}, sql string, args ...interface{}) error {
	return d.connection.Select(out, sql, args...)
}

// Executes "SELECT <x> WHERE <y> in (...)", result in "out" (casted)
func (d *PsqlDB) SelectIn(out interface{}, sql string, args ...interface{}) error {
	query, args, err := sqlx.In(sql, args...)
	if err != nil {
		return err
	}
	query = d.connection.Rebind(query)
	return d.connection.Select(out, query, args...)
}

// Executes "SELECT <x> WHERE <y> in (...)", result will be returned as sql rows
func (d *PsqlDB) QueryIn(sql string, args ...interface{}) (*sql.Rows, error) {
	query, args, err := sqlx.In(sql, args...)
	if err != nil {
		return nil, err
	}
	query = d.connection.Rebind(query)
	return d.connection.Query(query, args...)
}

// Executes query, usually non return values
func (d *PsqlDB) Exec(sql string, args ...interface{}) (sql.Result, error) {
	return d.connection.Exec(sql, args...)
}

// Executes query, usually have return values as sql rows
func (d *PsqlDB) ExecReturning(sql string, args ...interface{}) (*sql.Rows, error) {
	return d.connection.Query(sql, args...)
}
