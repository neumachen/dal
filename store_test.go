package dal

import (
	"database/sql"
	"log"
	"testing"

	"github.com/magicalbanana/vfs"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()
	db, openErr := sql.Open("postgres", "dbname=dbdal_test sslmode=disable")
	assert.NoError(t, openErr)

	fileStore, fsErr := vfs.LoadFiles("tests")
	assert.NoError(t, fsErr, "load filestore")

	dal := New(db, fileStore)
	assert.NotNil(t, dal)
}

func ExampleNew() {
	db, openErr := sql.Open("postgres", "dbname=dbdal_test sslmode=disable")
	if openErr != nil {
		log.Fatalln(openErr)
	}

	// vfs: github.com/magicalbanana/vfs
	fileStore, fsErr := vfs.LoadFiles("tests")
	if fsErr != nil {
		log.Fatalln(fsErr)
	}

	New(db, fileStore)
}

func TestOpen(t *testing.T) {
	t.Parallel()
	tests := []struct {
		desc      string
		driver    string
		dbcreds   func() string
		assertion func(t *testing.T, e error, m string)
	}{
		{
			// db connection fails because of the invalid driver being passed
			// to sqlx.Open
			desc:   "database fails to open",
			driver: "manbearpig",
			dbcreds: func() string {
				return "dbname=piggly_wiggly"
			},
			assertion: func(t *testing.T, e error, m string) {
				assert.Error(t, e, m)
			},
		},
		{
			desc:   "database loads successfully",
			driver: "postgres",
			dbcreds: func() string {
				return "dbname=piggly_wiggly"
			},
			assertion: func(t *testing.T, e error, m string) {
				assert.NoError(t, e, m)
			},
		},
	}

	for _, test := range tests {
		_, err := Open(test.driver, test.dbcreds())
		test.assertion(t, err, test.desc)
	}
}

func TestPingDataBase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		desc      string
		driver    string
		db        func() DB
		assertion func(t *testing.T, e error, m string)
	}{
		{
			// DB fails to ping because the database does not exist
			desc:   "database ping reaches maximum ping time",
			driver: "postgres",
			db: func() DB {
				db, _ := Open("postgres", "dbname=piggly_wiggly")
				return db
			},
			assertion: func(t *testing.T, e error, m string) {
				assert.Error(t, e, m)
			},
		},
		{
			desc:   "database loads successfully",
			driver: "postgres",
			db: func() DB {
				db, _ := Open("postgres", "dbname=dbdal_test sslmode=disable")
				return db
			},
			assertion: func(t *testing.T, e error, m string) {
				assert.NoError(t, e, m)
			},
		},
	}

	for _, test := range tests {
		lgr := func(msg string) {
			log.Println(msg)
		}
		pingErr := PingDatabase(test.db(), 2, lgr)
		test.assertion(t, pingErr, test.desc)
	}

}
