package pg_core_queries

import (
	"context"
	"errors"
	"freecreate/config"
	"freecreate/lib/api_error"
	"freecreate/lib/logger"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetUserByEmail(ctx context.Context, pgCoreQueries config.PgCoreQueries, pgxCore *pgxpool.Pool, email string) (int64, *api_error.Error) {

	query := pgCoreQueries.GetUserByEmail()
	queryArgs := pgx.NamedArgs{
		"email": email,
	}

	var userId int64
	queryErr := pgxCore.QueryRow(ctx, query, queryArgs).Scan(&userId)

	if errors.Is(queryErr, pgx.ErrNoRows) {
		apiErr := &api_error.Error{
			Code:    http.StatusNotFound,
			Message: "We could not find a user with that email address",
			Error:   queryErr,
		}

		return 0, apiErr
	} else if queryErr != nil {
		logger.Log(queryErr)

		apiErr := &api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   queryErr,
		}

		return 0, apiErr
	}

	return userId, nil
}
