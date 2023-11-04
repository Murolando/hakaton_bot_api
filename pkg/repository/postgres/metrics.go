package postgres

import "github.com/jmoiron/sqlx"

type MetricsPostgres struct {
	db *sqlx.DB
}

func NewMetricsPostgres(db *sqlx.DB) *MetricsPostgres {
	return &MetricsPostgres{
		db: db,
	}
}
