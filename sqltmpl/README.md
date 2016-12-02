# sqltmpl - Sequel Template

Allows the named parameters found in a SQL statement to be replaced with positional parameters.
It only currently supports Postgres

Example:

before:
```go
query := NewNamedParameterQuery("
    SELECT * FROM table
    WHERE col1 = :foo
    AND col2 IN(:firstName, :middleName, :lastName)
")

query.SetValue("foo", "bar")
query.SetValue("firstName", "Alice")
query.SetValue("lastName", "Bob")
query.SetValue("middleName", "Eve")

connection, _ := sql.Open("postgres", "user:pass@tcp(localhost:3306)/db")
connection.QueryRow(query.GetParsedQuery(), (query.GetParsedParameters())...)
```

after
```sql
SELECT * FROM boo AS b WHERE b.moo = $1
```

## USAGE
