package mocks

import (
	"database/sql"
	"errors"
)

// Db mocks out sqlx.DB for the purpose of testing.
type Db struct {
	PingOk     bool
	ExecOk     bool
	PrepareOk  bool
	QueryOk    bool
	QueryRowOk bool
}

// Prepare ...
func (db *Db) Prepare(query string) (*sql.Stmt, error) {
	if db.PrepareOk {
		return &sql.Stmt{}, nil
	}
	return nil, errors.New("mock error")
}

// Exec ...
func (db *Db) Exec(query string, args ...interface{}) (sql.Result, error) {
	if db.ExecOk {
		return &Result{}, nil
	}
	return nil, errors.New("mock error")
}

// Query ...
func (db *Db) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if db.QueryOk {
		return nil, nil
	}
	return nil, errors.New("mock error")
}

// QueryRow ...
func (db *Db) QueryRow(query string, args ...interface{}) *sql.Row {
	if db.QueryRowOk {
		return &sql.Row{}
	}
	return nil
}

// Ping ...
func (db *Db) Ping() error {
	if db.PingOk {
		return nil
	}
	return errors.New("mock Ping error")
}

// Result is a mock of sql.Result
type Result struct {
	LastInsertIDOk bool
}

// LastInsertId ...
func (r *Result) LastInsertId() (int64, error) {
	return 1, nil
}

// RowsAffected ...
func (r *Result) RowsAffected() (int64, error) {
	return 1, nil
}
