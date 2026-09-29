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

func GetMyCreatorWritings(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.GetMyCreatorWritingsParams) ([]pg_core_types.MyCreatorWritings, *api_error.Error) {
	var myCreatorWritings []pg_core_types.MyCreatorWritings

	query := pgCoreQueries.GetMyCreatorWritings()

	namedArgs := pgx.NamedArgs{
		"user_id":    params.UserId,
		"creator_id": params.CreatorId,
	}

	queryResult, queryErr := pgCore.Query(ctx, query, namedArgs)
	if queryErr != nil {
		logger.Log(queryErr)
		apiErr := &api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   queryErr,
		}

		return myCreatorWritings, apiErr
	}

	for queryResult.Next() {
		var creatorWriting pg_core_types.MyCreatorWritings

		scanErr := queryResult.Scan(&creatorWriting)
		if scanErr != nil {
			logger.Log(scanErr)
			apiErr := &api_error.Error{
				Code:    http.StatusInternalServerError,
				Message: api_error.InteralServerErrorMessage,
				Error:   scanErr,
			}

			return myCreatorWritings, apiErr
		}

		myCreatorWritings = append(myCreatorWritings, creatorWriting)
	}

	return myCreatorWritings, nil
}
