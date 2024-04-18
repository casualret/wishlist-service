package repo

import "server/internal/repo/postgres"

type Repos struct {
	db *postgres.Postgres
}

func InitRepo() (*Repos, error) {
	pg, err := postgres.InitPostgres()
	if err != nil {
		return nil, err
	}

	return &Repos{
		db: pg,
	}, nil
}
