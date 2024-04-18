package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sqlx.DB
}

func InitPostgres() (*Postgres, error) {
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=2281488 sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Postgres{
		db: db,
	}, nil
}
