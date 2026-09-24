package pg_core_queries

import (
	"context"
	"freecreate/internal/config"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetMyWritings(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.GetMyWritingsParams)([]pg_core_types.MyWritingsStruct, *api_error.Error){
	var myWritings []pg_core_types.MyWritingsStruct


	return myWritings, nil
}