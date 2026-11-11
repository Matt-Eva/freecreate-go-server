package pg_core_queries

import (
	"context"
	"freecreate/internal/config"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetUserUuid(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, userId int64)(uuid.UUID, *api_error.Error){
	var uuid uuid.UUID

	query := pgCoreQueries.GetUserUUID()

	namedArgs := pgx.NamedArgs{
		"user_id": userId,
	}

	queryErr := pgCore.QueryRow(ctx, query, namedArgs).Scan(&uuid)
	if queryErr != nil {
		logger.Log(queryErr)

		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: queryErr,
		}

		return uuid, apiErr
	}

	return uuid, nil
}