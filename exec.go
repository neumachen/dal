package dal

import (
	"database/sql"

	"github.com/magicalbanana/dal/sqltmpl"
)

func (d *dal) Exec(sqlFile string, params map[string]interface{}) (sql.Result, error) {
	qry, err := getSQLTemplate(sqlFile, d.fs)
	if err != nil {
		return nil, err
	}

	tmpl := sqltmpl.NewParser(qry)
	tmpl.SetValuesFromMap(params)

	stmt, err := d.db.Prepare(tmpl.GetParsedQuery())
	if err != nil {
		return nil, err
	}
	return stmt.Exec(tmpl.GetParsedParameters()...)
}
