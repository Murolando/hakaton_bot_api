package repository

import (
	"github.com/Murolando/hakaton_bot_api/pkg/repository/postgres"
	"github.com/jmoiron/sqlx"
)
type Metrics interface{

}
type Repository struct {
	Metrics
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Metrics: postgres.NewMetricsPostgres(db),
	}
}
