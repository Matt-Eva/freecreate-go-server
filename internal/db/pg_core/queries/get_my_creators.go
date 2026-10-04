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

func GetMyCreators(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.GetMyCreatorsParams) ([]pg_core_types.MyCreatorsStruct, *api_error.Error) {
	var myCreators []pg_core_types.MyCreatorsStruct

	query := pgCoreQueries.GetMyCreators()
	namedArgs := pgx.NamedArgs{
		"user_id": params.UserId,
	}

	queryResult, queryErr := pgCore.Query(ctx, query, namedArgs)

	if errors.Is(queryErr, pgx.ErrNoRows) {
		return myCreators, nil
	} else if queryErr != nil {
		logger.Log(queryErr)

		apiErr := &api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   queryErr,
		}
		return myCreators, apiErr
	}

	for queryResult.Next() {
		var myCreator pg_core_types.MyCreatorsStruct

		scanErr := queryResult.Scan(&myCreator)
		if scanErr != nil {
			logger.Log(scanErr)

			apiErr := &api_error.Error{
				Code:    http.StatusInternalServerError,
				Message: api_error.InteralServerErrorMessage,
				Error:   scanErr,
			}

			return myCreators, apiErr
		}

		myCreators = append(myCreators, myCreator)
	}

	return myCreators, nil
}
