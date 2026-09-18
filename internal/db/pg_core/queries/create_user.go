package pg_core_queries

import (
	"context"
	"errors"
	"freecreate/internal/config"
	pg_core_types "freecreate/internal/db/pg_core/types"
	pg_core_validators "freecreate/internal/db/pg_core/validators"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUser(ctx context.Context, pgCoreQueries config.PgCoreQueries, pgCore *pgxpool.Pool, email string) (pg_core_types.CreatedUser, *api_error.Error) {
	var createdUser pg_core_types.CreatedUser

	validationErr := pg_core_validators.ValidateEmail(email)
	if validationErr != nil {
		return createdUser, validationErr
	}

	query := pgCoreQueries.CreateUser()
	queryArgs := pgx.NamedArgs{
		"email": email,
	}

	var userId int64

	queryErr := pgCore.QueryRow(ctx, query, queryArgs).Scan(&userId)

	var pgErr *pgconn.PgError
	if errors.As(queryErr, &pgErr) && pgErr.Code == "23505" {
		apiErr := &api_error.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "An account with that email address already exists.",
			Error:   queryErr,
		}

		return createdUser, apiErr

	} else if queryErr != nil {
		logger.Log(queryErr)

		apiErr := &api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   queryErr,
		}

		return createdUser, apiErr
	}

	createdUser.UserId = userId

	return createdUser, nil
}
