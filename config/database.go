package config

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB(ctx context.Context, url string) *pgxpool.Pool {
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		panic(err)
	}

	return pool
}
