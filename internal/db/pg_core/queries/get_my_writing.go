package pg_core_queries

import (
	"context"
	"freecreate/internal/config"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetMyWriting(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.GetMyWritingParams)(pg_core_types.MyWriting, *api_error.Error) {
	var myWriting pg_core_types.MyWriting

	query := pgCoreQueries.GetMyWriting()

	namedArgs := pgx.NamedArgs{
		"user_id": params.UserId,
		"uuid": params.UUID,
	}

	queryErr := pgCore.QueryRow(ctx, query, namedArgs).Scan(&myWriting)
	if queryErr != nil {
		logger.Log(queryErr)
		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: queryErr,
		}
		
		return myWriting, apiErr
	}

	return myWriting, nil
}
