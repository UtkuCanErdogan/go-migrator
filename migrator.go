package go_migrator

import (
	"database/sql"
	"fmt"
)

type Migrator struct {
	Config Config
	db     *sql.DB
}

type Client int32

const (
	PostgresClient Client = 1
	MsSqlClient    Client = 2
)

type Config struct {
	Client     Client
	Connection Connection
	Migration  *Migration
	Seed       *Seed
}

type Connection struct {
	Username string
	Password string
	Host     string
	Port     int32
	Database string
}

type Migration struct {
	TableName *string
	Directory *string
}

type Seed struct {
	TableName *string
	Directory *string
}

func New(config Config) (*Migrator, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Connection.Host, config.Connection.Port, config.Connection.Username, config.Connection.Password, config.Connection.Database)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Migrator{Config: config, db: db}, nil
}

func (m *Migrator) getDbInstance() *sql.DB {
	return m.db
}
