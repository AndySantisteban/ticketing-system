package connection

import (
	"InfositelOR/pkg/infrastructure/db/mysql/persistence"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	// connectionString = "postgres://andysantisteban:rbMeQeXFAa5weA0GZLFcp2pPZRkevvsJ@dpg-cokjc70l5elc73c7phvg-a.oregon-postgres.render.com/infositelordb"
	connectionString = "postgres://postgres:pasword@127.0.0.1/infositelordb"
)

func NewConnection() (*sql.DB, *persistence.Queries, error) {

	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to ping database: %w", err)
	}
	// defer database.Close()

	err = database.Ping()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to ping database: %w", err)
	}

	queries := persistence.New(database)

	return database, queries, nil
}
