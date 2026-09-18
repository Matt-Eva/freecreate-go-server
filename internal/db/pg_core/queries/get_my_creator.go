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

func GetMyCreator(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, creatorUuid string, userId int64) (pg_core_types.MyCreator, *api_error.Error) {
	query := pgCoreQueries.GetMyCreator()

	namedArgs := pgx.NamedArgs{
		"user_id": userId,
		"uuid":    creatorUuid,
	}

	var myCreator pg_core_types.MyCreator

	queryErr := pgCore.QueryRow(ctx, query, namedArgs).Scan(&myCreator)
	if errors.Is(queryErr, pgx.ErrNoRows) {
		apiErr := &api_error.Error{
			Code:    http.StatusNotFound,
			Message: "Hmm, we couldn't find the creator you're looking for...",
			Error:   queryErr,
		}

		return myCreator, apiErr
	} else if queryErr != nil {
		logger.Log(queryErr)

		apiErr := &api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   queryErr,
		}

		return myCreator, apiErr
	}

	return myCreator, nil
}
