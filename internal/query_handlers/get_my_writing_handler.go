package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"

	"github.com/jackc/pgx/v5/pgxpool"
)

func HandleGetMyWriting(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.GetMyWritingParams)(pg_core_types.MyWriting, *api_error.Error) {
	myWriting, queryErr := pg_core_queries.GetMyWriting(ctx, pgCore, pgCoreQueries, params)
	
	return myWriting, queryErr
}
