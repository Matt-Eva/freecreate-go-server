package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"

	"github.com/jackc/pgx/v5/pgxpool"
)

func HandleUpdateUserInfo(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.UpdateUserInfoParams) (pg_core_types.UserInfo, *api_error.Error) {

	info, getInfoErr := pg_core_queries.UpdateUserInfo(ctx, pgCore, pgCoreQueries, params)
	if getInfoErr != nil {
		return info, getInfoErr
	}

	return info, nil
}
