package pg_core_queries

import (
	"context"
	"errors"
	"freecreate/config"
	"freecreate/lib/logger"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetUserByEmail(ctx context.Context, pgCoreQueries config.PgCoreQueries, pgxCore *pgxpool.Pool, email string) (int, error) {

	query := pgCoreQueries.GetUserByEmail()
	queryArgs := pgx.NamedArgs{
		"email": email,
	}

	var userId int
	queryErr := pgxCore.QueryRow(ctx, query, queryArgs).Scan(&userId)
	if errors.Is(queryErr, pgx.ErrNoRows) {
		return 0, queryErr
	} else if queryErr != nil {
		logger.Log(queryErr)
		return 0, queryErr
	}

	return userId, nil
}
