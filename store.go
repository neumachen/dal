package dal

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	// postgres driver
	_ "github.com/lib/pq"
)

// DAL (database access layer) is the type returned when initializing an new
// instance of this package. The functions listed are the functions that can
// be used for this package.
type DAL interface {
	Exec(sqlFile string, params map[string]interface{}) (sql.Result, error)
	Query(sqlFile string, params map[string]interface{}) (*sql.Rows, error)
	QueryRow(sqlFile string, params map[string]interface{}) (*sql.Row, error)
}

// DB is an interface modeled after the go's standard database/sql package.
type DB interface {
	Prepare(query string) (*sql.Stmt, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Ping() error
}

// FileStore is a package agnostic interface for any type of virtual file
// system that returns the value as string.
//
// Of course the string MUST be the SQL template that will be used upon
// executing a query function (Query, QueryRow, etc).
type FileStore interface {
	Get(file string) (string, error)
}

type dal struct {
	db DB
	fs FileStore
}

// New returns a DAL using the given DB and FileStore.
func New(db DB, fileStore FileStore) DAL {
	return &dal{
		db: db,
		fs: fileStore,
	}
}

// Open opens a new database connection using the given driver and the
// dbCreds. If an error occurs when opening a connection an error is
// returned.
func Open(driver string, dbCreds string) (DB, error) {
	db, err := sql.Open(driver, dbCreds)
	if err != nil {
		return nil, fmt.Errorf("database connection failed")
	}
	return db, nil
}

// PingDatabase is a helper function to ping the database with backoff
// to ensure a connection can be established before we proceed with a
// database setup whilst logging each ping and returns an error if it fails
// using the given function to log.
func PingDatabase(db DB, dbPingTime int, lgr func(msg string)) error {
	for i := 0; i < dbPingTime; i++ {
		err := db.Ping()
		if err == nil {
			return err
		}
		if lgr != nil {
			lgr("Database ping failed. Retry in 2s")
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("failed to ping the database after %v seconds", dbPingTime)
}

// getSQLTemplate gets the given sqlFile from the virtual file system and
// returns the sqlFile's content. If virtual file system does not hold the
// SQL file an error is returned while returning the zero value for the for
// the sqlStatemetn.
func getSQLTemplate(sqlFile string, fs FileStore) (string, error) {
	if fs == nil {
		return "", errors.New("empty FileStore")
	}

	sqlStatement, err := fs.Get(sqlFile)

	if err != nil {
		return "", err

	}
	return sqlStatement, nil
}
