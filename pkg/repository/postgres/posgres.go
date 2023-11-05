package postgres

import (
	"fmt"

	"github.com/Murolando/hakaton_bot_api/ent"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	userTable = "user"
)


func NewPostgresDB(cfg *ent.Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host = %s port = %s user = %s dbname = %s password = %s  sslmode = disable",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
