package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"InfositelOR/pkg/infrastructure/db/mysql/persistence"
)

const (
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
