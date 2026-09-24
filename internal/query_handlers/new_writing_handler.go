package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"

	"github.com/jackc/pgx/v5/pgxpool"
)

func HandleNewWriting(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.NewWritingParams) (pg_core_types.CreatedWriting, *api_error.Error) {
	createdWriting, createWritingErr := pg_core_queries.CreateWriting(ctx, pgCore, pgCoreQueries, params)
	return createdWriting, createWritingErr
}
