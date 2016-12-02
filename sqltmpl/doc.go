// Package sqltmpl provides support for named parameters in SQL queries
// used by Go / golang programs and libraries.
//
// Named parameters are not supported by all SQL query engines, and their standards are scattered.
// But positional parameters have wide adoption across all databases.
//
// npq package translates SQL queries which use named parameters into queries which use positional parameters.
//
// Example usage:
//
// 	query := NewNamedParameterQuery("
// 		SELECT * FROM table
// 		WHERE col1 = :foo
// 	")
//
// 	query.SetValue("foo", "bar")
//
// 	connection, _ := sql.Open("mysql", "user:pass@tcp(localhost:3306)/db")
// 	connection.QueryRow(query.GetParsedQuery(), (query.GetParsedParameters())...)
//
// In the example above, note the format of "QueryRow". In order to use named parameter queries,
// you will need to use npq exact format, including the variadic symbol "..."
//
// Note that the example above uses "QueryRow", but named parameters used in npq fashion
// work equally well for "Query" and "Exec".
//
// It's also possible to pass in a map, instead of defining each parameter individually:
//
// 	query := NewNamedParameterQuery("
// 		SELECT * FROM table
// 		WHERE col1 = :foo
// 		AND col2 IN(:firstName, :middleName, :lastName)
// 	")
//
// 	var parameterMap = map[string]interface{} {
// 		"foo": 		"bar",
// 		"firstName": 	"Alice",
// 		"lastName": 	"Bob"
// 		"middleName": 	"Eve",
// 	}
//
// 	query.SetValuesFromMap(parameterMap)
//
// 	connection, _ := sql.Open("mysql", "user:pass@tcp(localhost:3306)/db")
// 	connection.QueryRow(query.GetParsedQuery(), (query.GetParsedParameters())...)
package sqltmpl
