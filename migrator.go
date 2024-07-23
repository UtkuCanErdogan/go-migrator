package go_migrator

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type TableBuilder struct {
	createBuilder *CreateTableBuilder
	alterBuilder  *AlterTableBuilder
}
type Migrate struct {
	key          string
	fn           func(migrator *Migrator) *TableBuilder
	TableBuilder *TableBuilder
}

type Migrator struct {
	Config  Config
	migrate []Migrate
	db      *sql.DB
}

type Client int32

const (
	PostgresClient Client = 1
)

type Config struct {
	Client     Client
	Connection *Connection
	Gorm       *gorm.DB
	Schema     *string
}

type Connection struct {
	Username string
	Password string
	Host     string
	Port     int32
	Database string
}

func New(config Config) (*Migrator, error) {
	if config.Connection == nil && config.Gorm == nil {
		return nil, errors.New("no connection parameter found")
	}

	var db *sql.DB
	var err error
	if config.Connection != nil {
		connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Connection.Host, config.Connection.Port, config.Connection.Username, config.Connection.Password, config.Connection.Database)
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			return nil, err
		}
	} else if config.Gorm != nil {
		db, err = config.Gorm.DB()
		if err != nil {
			return nil, err
		}
	}

	if config.Schema == nil {
		schema := "public"
		config.Schema = &schema
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Migrator{Config: config, db: db}, nil
}

func (m *Migrator) getDbInstance() *sql.DB {
	return m.db
}

func (m *Migrator) AddMigration(key string, fn func(migrator *Migrator) *TableBuilder) {
	migrate := Migrate{key: key, fn: fn}
	m.migrate = append(m.migrate, migrate)
}
