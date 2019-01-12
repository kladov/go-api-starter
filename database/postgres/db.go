package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// New return connection to postgres
// connnStr example: "postgres://user:password@localhost/database?sslmode=disable"
// more information https://godoc.org/github.com/lib/pq
func New(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
