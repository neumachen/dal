// Package dal is an opinionated package that wraps the functions of  go's
// standard database/sql package. The versions of sql.Query, sql.QueryRow, et
// al. all are left untouched so that the interfaces are a superset of the
// standard ones.
//
// The key opinionated approaches used by this libraray is the implementation
// of an interface: type FileStore interface { Get(sqlTemplate string)
// (string, error) } that MUST return the string of a SQL template. The other
// opinionated approach is that the parameters passed on the the Query,
// QueryRow functions be of: map[string]interface{} versus: interface{} This
// makes the execution a bit faster than other libraries.
package dal
