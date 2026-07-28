package pg_core_queries

import (
	"context"
	"errors"
	"freecreate/config"
	pg_core_validators "freecreate/db/pg_core/validators"
	"freecreate/lib/logger"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUser(ctx context.Context, pgCoreQueries config.PgCoreQueries, pgCore *pgxpool.Pool, email string) (int64, error) {
	validationErr := pg_core_validators.ValidateEmail(email)
	if validationErr != nil {
		logger.Log(validationErr)
		return 0, validationErr
	}
	query := pgCoreQueries.CreateUser()
	queryArgs := pgx.NamedArgs{
		"email": email,
	}

	var userId int64

	queryErr := pgCore.QueryRow(ctx, query, queryArgs).Scan(&userId)
	if queryErr != nil {
		logger.Log(queryErr)
		err := errors.New("could not create user")
		return 0, err
	}
	return userId, nil
}
