package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"

	"github.com/jackc/pgx/v5/pgxpool"
)

func HandleGetMyCreator(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.GetMyCreatorParams) (pg_core_types.MyCreator, *api_error.Error) {

	myCreator, err := pg_core_queries.GetMyCreator(ctx, pgCore, pgCoreQueries, params)

	return myCreator, err
}
