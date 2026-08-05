package pg_core_queries

import (
	"context"
	"freecreate/config"
	pg_core_validators "freecreate/db/pg_core/validators"
	"freecreate/lib/api_error"
	"freecreate/lib/logger"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUser(ctx context.Context, pgCoreQueries config.PgCoreQueries, pgCore *pgxpool.Pool, email string) (int64, *api_error.Error) {
	validationErr := pg_core_validators.ValidateEmail(email)
	if validationErr != nil {
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

		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: queryErr,
		}

		return 0, apiErr
	}
	return userId, nil
}
