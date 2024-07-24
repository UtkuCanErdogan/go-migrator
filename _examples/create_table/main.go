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
		Schema: nil,
		Gorm:   nil,
	}

	m, err := migrator.New(config)
	if err != nil {
		log.Fatal(err)
	}

	m.AddMigration("RoleCreate", roleCreate)
	m.AddMigration("AccountCreate", accountCreate)

	if err = m.Migrate(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration succeed")
}

func accountCreate(migrate *migrator.Migrator) *migrator.TableBuilder {
	account := migrate.CreateTable("account")
	account.UUID("id").Primary()
	account.String("first_name", 50).NotNull()
	account.String("last_name", 50).NotNull()
	account.String("email", 80).Unique().NotNull()
	account.String("phone_number", 16).Unique().NotNull()
	account.Boolean("is_active").NotNull().Default(true)
	account.UUID("role_id").Foreign().Table("role").Reference("id")

	return account.Build()
}

func roleCreate(migrate *migrator.Migrator) *migrator.TableBuilder {
	role := migrate.CreateTable("role")
	role.UUID("id").Primary()
	role.String("first_name", 50).NotNull()

	return role.Build()
}
