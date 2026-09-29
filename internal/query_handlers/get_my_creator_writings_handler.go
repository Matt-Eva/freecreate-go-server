package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"

	"github.com/jackc/pgx/v5/pgxpool"
)

func HandleGetMyCreatorWritings(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.GetMyCreatorWritingsParams) ([]pg_core_types.MyCreatorWritings, *api_error.Error) {

	myCreatorWritings, err := pg_core_queries.GetMyCreatorWritings(ctx, pgCore, pgCoreQueries, params)

	return myCreatorWritings, err
}
