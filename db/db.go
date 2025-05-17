package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type Postgres struct {
	Db *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
	ConnString string
)

func NewPG(ctx context.Context) (*Postgres, error) {
	var pgErr error = nil
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, ConnString)
		if err != nil {
			pgErr = fmt.Errorf("unable to create connection pool: %w", err)
			return
		}

		pgInstance = &Postgres{db}
	})

	if pgErr != nil {
		return nil, pgErr
	}

	return pgInstance, nil
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.Db.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.Db.Close()
}
