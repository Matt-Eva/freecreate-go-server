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

func GetMyWritings(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.GetMyWritingsParams) ([]pg_core_types.MyWritingsStruct, *api_error.Error) {
	var myWritings []pg_core_types.MyWritingsStruct

	query := pgCoreQueries.GetMyWritings()

	namedArgs := pgx.NamedArgs{
		"user_id": params.UserId,
	}

	queryResult, queryErr := pgCore.Query(ctx, query, namedArgs)
	if queryErr != nil {
		logger.Log(queryErr)
		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: queryErr,
		}

		return myWritings, apiErr
	}

	for queryResult.Next(){
		var writing pg_core_types.MyWritingsStruct

		scanErr := queryResult.Scan(&writing)
		if scanErr != nil {
			logger.Log(scanErr)
			apiErr := &api_error.Error{
				Code: http.StatusInternalServerError,
				Message: api_error.InteralServerErrorMessage,
				Error: scanErr,
			}

			return myWritings, apiErr
		}

		myWritings = append(myWritings, writing)
	}

	return myWritings, nil
}
