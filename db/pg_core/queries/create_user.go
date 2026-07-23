package pg_core_queries

import (
	"context"
	"freecreate/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUser(ctx context.Context, pgCoreQueries config.PgCoreQueries, pgCore *pgxpool.Pool, email string) (int, error) {
	return 0, nil
}
