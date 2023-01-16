package config

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

const (
	Hostname = "localhost"
	Port     = "5432"
	Username = "postgres"
	Password = "postgres"
	Dbname   = "go_restful"
)

func InitDBTest(ctx context.Context, url string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, url)

	return pool, err
}

func TestConnectionSuccess(t *testing.T) {
	t.Run("Test Scenario DB Connection Success", func(t *testing.T) {
		ctx := context.Background()
		url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", Username, Password, Hostname, Port, Dbname)

		pool, err := InitDBTest(ctx, url)
		defer pool.Close()

		err = pool.Ping(ctx)

		assert.Nil(t, err)
		assert.NotNil(t, pool)
	})
}

func TestConnectionFailed(t *testing.T) {
	t.Run("Test Scenario DB Connection Failed", func(t *testing.T) {
		ctx := context.Background()
		url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", "must_failed", "Password", Hostname, Port, "123")

		pool, err := InitDBTest(ctx, url)
		defer pool.Close()

		err = pool.Ping(ctx)

		assert.NotNil(t, err)
		assert.NotNil(t, pool)
	})
}
