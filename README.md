
[Build Status]: https://travis-ci.org/magicalbanana/dal
[Build Status Badge]: https://travis-ci.org/magicalbanana/dal.svg?branch=master

[Coverage Status]: https://coveralls.io/github/magicalbanana/dal?branch=master
[Coverage Status Badge]: https://coveralls.io/repos/github/magicalbanana/dal/badge.svg?branch=master

[Doc]: https://godoc.org/github.com/magicalbanana/dal
[Doc Badge]: https://godoc.org/github.com/magicalbanana/dal?status.svg

[Go Report Card]: https://goreportcard.com/report/github.com/magicalbanana/dal
[Go Report Card Badge]: https://goreportcard.com/badge/github.com/magicalbanana/dal

[![Build Status][Build Status Badge]][Build Status]
[![Coverage Status][Coverage Status Badge]][Coverage Status]
[![Doc][Doc Badge]][Doc]
[![Go Report Card][Go Report Card Badge]][Go Report Card]

# dal - Database Access Layer

## Description

This package is very opinionated. It aims to eliminate inline code SQL statemnts by requiring a virtual file system that will hold all the sql templates in memory during runtime.

This package wraps the `database/sql` package.

## Development

This package is still under active development. It is wise to vendor this package because although not planned some breaking API changes may be introduced.

## Required Interfaces

```go
type FileStore interface {
    Get(file string) (string, error)
}
```

The interface `FileStore` is a required interface that must be implemented in order to instantiate a new `DAL` because this is need when calling the query methods for the purpose of parsing the SQL template.

## Usage

```go
package main

func main() {
    lgr := func(msg string) {
        log.Println(msg)
    }

    dataStore, err := loadDataStore(envcfg.PgDbCreds(), envcfg.DbPingTime(), fileStore, lgr)
    if err != nil {
        log.Fatalln(err)
    }
}


func loadDataStore(dbCreds string, dbPingTime int, fileStore vfs.Store, lgr finlog.Logger) (dal.DAL, error) {
    pgdal, openErr := dal.Open("postgres", envcfg.PgDbCreds())
    if openErr != nil {
            return nil, openErr
    }

    lgrFunc := func(msg string) {
            lgr.Info(msg)
    }

    pingErr := dal.PingDatabase(pgdal, dbPingTime, lgrFunc)
    if pingErr != nil {
            return nil, pingErr
    }

    return dal.New(pgdal, fileStore), nil
}
```
