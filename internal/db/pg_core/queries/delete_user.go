package pg_core_queries

import (
	"context"
	"errors"
	"freecreate/internal/config"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func DeleteUser(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.DeleteUserParams) *api_error.Error {
	namedArgs := pgx.NamedArgs{
		"id": params.UserId,
	}

	query := pgCoreQueries.DeleteUser()

	output, err := pgCore.Exec(ctx, query, namedArgs)
	if err != nil {
		logger.Log(err)

		apiErr := &api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   err,
		}

		return apiErr
	} else if output.RowsAffected() == 0 {
		msg := "there were no matching rows to delete"
		err := errors.New(msg)
		logger.Log(err)

		apiErr := &api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: msg,
			Error:   err,
		}

		return apiErr
	}

	return nil
}
