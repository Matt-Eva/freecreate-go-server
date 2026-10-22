package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"

	"github.com/jackc/pgx/v5/pgxpool"
)

func HandleUpdateChapterNumber(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.UpdateChapterNumberParams) (int, *api_error.Error){

	newChapterNumber, err := pg_core_queries.UpdateChapterNumber(ctx, pgCore, pgCoreQueries, params)

	return newChapterNumber, err
}
