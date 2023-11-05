package postgres

import (
	"fmt"

	"github.com/Murolando/hakaton_bot_api/ent"
	"github.com/jmoiron/sqlx"
)

type MetricsPostgres struct {
	db *sqlx.DB
}

func NewMetricsPostgres(db *sqlx.DB) *MetricsPostgres {
	return &MetricsPostgres{
		db: db,
	}
}

func (r *MetricsPostgres) DataBaseStat() (*ent.DataBaseStat, error) {
	var resp ent.DataBaseStat
	query := fmt.Sprintf(`
	SELECT sum(active_time),  
		sum(sessions), 
		sum(tup_inserted), 
		sum(tup_updated), 
		sum(tup_deleted)  
	FROM pg_stat_database`)
	row := r.db.QueryRow(query)
	if err := row.Scan(&resp.SumActiveTime,&resp.SumSessions,&resp.SumInsertedDatas,
		&resp.SumUpdatedDatas,&resp.SumDeletedDatas); err != nil {
		return nil, err
	}
	query = fmt.Sprintf(`
	SELECT sum(xact_commit), sum(xact_rollback)  
	FROM pg_stat_database`)
	row = r.db.QueryRow(query)
	if err := row.Scan(&resp.CommitNum,&resp.RollabackNum); err != nil {
		return nil, err
	}
	return &resp, nil
}
func (r *MetricsPostgres) PgActivityStat() (*ent.PgActivityStat, error) {
	var resp ent.PgActivityStat
	query := fmt.Sprintf(`SELECT count(*) FROM pg_stat_activity WHERE state = 'active'`)
	row := r.db.QueryRow(query)
	if err := row.Scan(&resp.ActualSessions); err != nil {
		return nil, err
	}

	query = fmt.Sprintf(`SELECT count(*) FROM pg_stat_activity`)
	row = r.db.QueryRow(query)
	if err := row.Scan(&resp.AllSessions); err != nil {
		return nil, err
	}
	return &resp, nil
}
func (r *MetricsPostgres) PgDatabaseSize(bdName string) (*ent.PgDatabaseSize, error) {
	var resp ent.PgDatabaseSize
	query := fmt.Sprintf(`SELECT pg_database_size($1);`)
	row := r.db.QueryRow(query,bdName)
	if err := row.Scan(&resp.CurrentMem); err != nil {
		return nil, err
	}
	return &resp, nil
}
