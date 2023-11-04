package repository

import (
	"github.com/Murolando/default_rest_arch/ent"
	"github.com/Murolando/default_rest_arch/pkg/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	SignUp(user ent.User) (int64, error)
	GetUserByLoginAndPassword(mail *string, password *string) (int64, error)
	SetSession(user int64, refresh string, expiredAt string) error
	GetByRefreshToken(refresh string) (int64, error)
}

type Repository struct {
	Auth
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: postgres.NewAuthPostgres(db),
	}
}
