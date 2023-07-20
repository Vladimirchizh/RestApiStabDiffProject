package db

import (
	"RestApiStabDiffProject/internal/queries"
	"database/sql"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err

	}
	return &PostgresStore{db: db}, nil

}

func (s *PostgresStore) CreateQueryTable() (string, error) {
	return "", nil
}

func (s *PostgresStore) Create(q queries.Query) error {
	return nil
}
