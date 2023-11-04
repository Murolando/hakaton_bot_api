package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

const (
	userTable = "user"
)

func NewConfig(host string, port string, user string, password string, dbname string) *Config {
	return &Config{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbname:   dbname,
	}
}

func NewPostgresDB(cfg *Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host = %s port = %s user = %s dbname = %s password = %s  sslmode = disable",
		cfg.host, cfg.port, cfg.user, cfg.dbname, cfg.password)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
