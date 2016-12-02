package dal

import (
	"database/sql"
	"testing"

	"github.com/magicalbanana/vfs"
	"github.com/stretchr/testify/assert"
)

func TestRowsToMap(t *testing.T) {
	t.Parallel()
	tests := []struct {
		desc      string
		assertion func(t *testing.T, desc string)
	}{
		{
			desc: "maps row columns to a map[string]interface{}",
			assertion: func(t *testing.T, desc string) {
				// arrangement
				fs, fsErr := vfs.LoadFiles("tests/fixtures/sqls")
				assert.NoError(t, fsErr)
				db, openErr := sql.Open("postgres", "dbname=dbdal_test sslmode=disable")
				assert.NoError(t, openErr)
				params := map[string]interface{}{
					"first_name": "piggly",
					"last_name":  "wiggly",
					"address":    []byte(`{"test": "pigglywiggly"}`),
				}
				d := &dal{db: db, fs: fs}
				_, qryErr := d.Query("insert_customer.sql", params)
				assert.NoError(t, qryErr, "insert customer")
				rows, qryErr := d.Query("select_customer.sql", params)
				assert.NoError(t, qryErr)

				// act
				m := RowsToMap(rows)

				// assertion
				assert.Equal(t, m[0]["first_name"], "piggly", desc)
				assert.Equal(t, m[0]["last_name"], "wiggly", desc)

				// clean up
				_, qryErr = d.Query("delete_customer.sql", params)
				assert.NoError(t, qryErr, "delete customer")
			},
		},
	}

	for _, test := range tests {
		test.assertion(t, test.desc)
	}
}
