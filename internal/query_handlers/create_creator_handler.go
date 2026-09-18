package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"

	"github.com/jackc/pgx/v5/pgxpool"
)

func HandleCreateCreator(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, queryParams pg_core_types.NewCreatorParams) (pg_core_types.CreatedCreator, *api_error.Error) {

	creator, createCreatorErr := pg_core_queries.CreateCreator(ctx, pgCore, pgCoreQueries, queryParams)
	if createCreatorErr != nil {
		return creator, createCreatorErr
	}

	return creator, nil
}
