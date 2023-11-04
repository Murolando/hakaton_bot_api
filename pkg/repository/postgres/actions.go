package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ActionPostgres struct {
	db *sqlx.DB
}

func NewActionPostgres(db *sqlx.DB) *ActionPostgres {
	return &ActionPostgres{
		db: db,
	}
}

// lock table
func (r *ActionPostgres) StopBase(tableName string) (bool, error) {
	var result bool
	query := fmt.Sprintf(`SELECT * FROM "%s" FOR UPDATE`,tableName)
	row := r.db.QueryRow(query)
	if err := row.Scan(&result); err != nil {
		return false, err
	}
	return result, nil
}

func (r *ActionPostgres) RunBase(tableName string) (bool, error) {
	var result bool
	query := fmt.Sprintf(`UNLOCK TABLE "%s"`,tableName)
	row := r.db.QueryRow(query)
	if err := row.Scan(&result); err != nil {
		return false, err
	}
	return result, nil
}

func (r *ActionPostgres) LastBaseCheckpoint() error{
	query := fmt.Sprintf(`CHECKPOINT`)
	_,err := r.db.Exec(query)
	if err!=nil{
		return err
	}
	return nil
}

func (r *ActionPostgres) KillProcess(pid int64) (bool, error) {
	var result bool
	query := fmt.Sprintf(`SELECT pg_terminate_backend($1)`)
	row := r.db.QueryRow(query,pid)
	if err := row.Scan(&result); err != nil {
		return false, err
	}
	return result, nil
}
func (r *ActionPostgres) UserAction() {

}