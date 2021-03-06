package store

import (
	"database/sql"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

type storePostgres struct {
	db *sql.DB
}

func NewPostgresMigrator(connectionString string) (Migrator, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return &storePostgres{
		db: db,
	}, nil
}

func NewPostgresStore(connectionString string) (Store, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return &storePostgres{
		db: db,
	}, nil
}

func (store *storePostgres) MigrateDB() error {
	migrations := allMigrations()
	n, err := migrate.Exec(store.db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}
	fmt.Println("db.migrations: applied %d migrations", n)

	return nil
}

func (store *storePostgres) ReverseDB() error {
	migrations := allMigrations()
	n, err := migrate.Exec(store.db, "postgres", migrations, migrate.Down)
	if err != nil {
		return err
	}
	fmt.Println("db.migrations: reversed %d migrations", n)

	return nil
}

func allMigrations() *migrate.MemoryMigrationSource {
	return &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			{
				Id: "1",
				Up: []string{`
			CREATE TABLE users (
				id SERIAL PRIMARY KEY,
				name  VARCHAR NOT NULL,
				created_at TIMESTAMPTZ NOT NULL,
				updated_at TIMESTAMPTZ NOT NULL
			)`},
				Down: []string{`DROP TABLE users`},
			},
		},
	}
}
