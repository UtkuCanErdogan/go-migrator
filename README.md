# Go Migrator

The migration library for Golang, aims to be developer friendly

# Installing
Using Migrator is easy. First, use `go get` to install the latest version
of the library.

```
go get -u github.com/UtkuCanErdogan/go-migrator
```

Next, include Migrator in your application:

```go
import "github.com/UtkuCanErdogan/go-migrator"
```

# Overview
Go Migrator is a library providing perform migration operations with the Migration builder. Inspired by [Knex.js](https://knexjs.org/)

Migrator provides:
* Full support of POSTGRESQL
* Create, Alter operations
* Runs each migration once