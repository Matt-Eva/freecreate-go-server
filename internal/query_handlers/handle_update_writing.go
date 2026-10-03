package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"

	"github.com/jackc/pgx/v5/pgxpool"
)

func HandleUpdateWriting(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.UpdateWritingParams) (pg_core_types.UpdatedWriting, *api_error.Error) {

	updatedWriting, updateWritingErr := pg_core_queries.UpdateWriting(ctx, pgCore, pgCoreQueries, params)

	return updatedWriting, updateWritingErr
}
