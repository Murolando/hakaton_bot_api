package repository

import (
	"github.com/Murolando/hakaton_bot_api/ent"
	"github.com/Murolando/hakaton_bot_api/pkg/repository/postgres"
	"github.com/jmoiron/sqlx"
)
type Actions interface{
	StopBase(tableName string) (bool, error)
	RunBase(tableName string) (bool, error)

	LastBaseCheckpoint() error
	KillProcess(pid int64) (bool, error)
}
type Metrics interface{
	DataBaseStat()(*ent.DataBaseStat,error) 
	PgActivityStat()(*ent.PgActivityStat,error)
	PgDatabaseSize(bdName string) (*ent.PgDatabaseSize, error)
}
type Repository struct {
	Actions
	Metrics
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Metrics: postgres.NewMetricsPostgres(db),
		Actions: postgres.NewActionPostgres(db),
	}
}
