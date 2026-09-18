package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"

	"github.com/jackc/pgx/v5/pgxpool"
)

func HandleGetUserInfo(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, userId int64) (pg_core_types.UserInfo, *api_error.Error) {

	userInfo, getUserErr := pg_core_queries.GetUserInfo(ctx, pgCore, pgCoreQueries, userId)
	if getUserErr != nil {
		return userInfo, getUserErr
	}

	return userInfo, nil
}
