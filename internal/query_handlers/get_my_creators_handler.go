package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"

	"github.com/jackc/pgx/v5/pgxpool"
)

func HandleGetMyCreators(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.GetMyCreatorsParams) ([]pg_core_types.MyCreatorsStruct, *api_error.Error) {

	myCreators, getMyCreatorsErr := pg_core_queries.GetMyCreators(ctx, pgCore, pgCoreQueries, params)
	if getMyCreatorsErr != nil {
		return myCreators, getMyCreatorsErr
	}

	return myCreators, nil
}
