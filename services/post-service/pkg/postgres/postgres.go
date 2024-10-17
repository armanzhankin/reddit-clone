package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPgxPool(ctx context.Context, dbUrl string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
