package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	var migrationPath, migrationTable, DBurl string
	flag.StringVar(&migrationPath, "migrations-path", "", "Path to migration folder")
	flag.StringVar(&migrationTable, "migrations-table", "migrations", "Name of the migration table")
	flag.StringVar(&DBurl, "db-url", "", "Database connection string")
	flag.Parse()

	if migrationPath == "" {
		panic("Missing migrations")
	}

	m, err := migrate.New("file://"+migrationPath, fmt.Sprintf("postgres://%s?x-migrations-table=%s&sslmode=disable", DBurl, migrationTable))
	if err != nil {
		panic(fmt.Errorf("failed to create migration instance: %s", err))
	}

	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")
			return
		}
		panic(fmt.Errorf("migration failed: %w", err))
	}

	fmt.Println("migrations applied successfully")
}
