package pgx

import (
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Repo struct {
	pool  *pgxpool.Pool
	log   *zap.Logger
	pgErr *pgconn.PgError
}

func New(pool *pgxpool.Pool, log *zap.Logger, pgErr *pgconn.PgError) (*Repo, error) {
	repo := &Repo{
		pool:  pool,
		log:   log,
		pgErr: pgErr,
	}

	return repo, nil
}
