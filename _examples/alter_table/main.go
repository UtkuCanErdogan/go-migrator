package main

import (
	"fmt"
	migrator "github.com/UtkuCanErdogan/go-migrator"
	"log"
)

func main() {
	config := migrator.Config{
		Client: migrator.PostgresClient,

		Connection: &migrator.Connection{
			Database: "migrator",
			Username: "postgres",
			Password: "password",
			Host:     "localhost",
			Port:     5432,
		},
		Gorm: nil,
	}

	m, err := migrator.New(config)
	if err != nil {
		log.Fatal(err)
	}

	m.AddMigration("AccountAlter", accountAlter)

	if err = m.Migrate(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration succeed")
}

func accountAlter(migrate *migrator.Migrator) *migrator.TableBuilder {
	account := migrate.AlterTable("account")
	account.DropConstraint("email", "email_unique", migrator.UniqueConstraint)
	account.DropConstraint("phone_number", "phone_number_unique", migrator.UniqueConstraint)
	account.DropColumn("is_active")
	account.Integer("status")
	account.RenameColumn("phone_number", "mobile_phone")

	return account.Build()
}
