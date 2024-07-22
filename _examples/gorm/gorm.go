package main

import (
	"fmt"
	migrator "github.com/UtkuCanErdogan/go-migrator"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:password@localhost:5432/migrator"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	config := migrator.Config{
		Client:     migrator.PostgresClient,
		Connection: nil,
		Gorm:       db,
	}

	m, err := migrator.New(config)
	if err != nil {
		log.Fatal(err)
	}

	m.AddMigration("CreateTable", createTable)
	if err = m.Migrate(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration succeed")
}

func createTable(migrate *migrator.Migrator) *migrator.TableBuilder {
	account := migrate.CreateTable("account")
	account.UUID("id").Primary()
	account.String("first_name", 50).NotNull()
	account.String("last_name", 50).NotNull()
	account.String("email", 80).Unique().NotNull()

	return account.Build()
}
