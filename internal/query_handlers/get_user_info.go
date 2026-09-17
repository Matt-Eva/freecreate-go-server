package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	"freecreate/internal/lib/api_error"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserInfo struct {
	Username       string
	UserHandle     string
	IsAdult        bool
	ReadingHistory bool
}

func HandleGetUserInfo(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, userId int64)(UserInfo, *api_error.Error) {
	var userInfo UserInfo

	user, getUserErr := pg_core_queries.GetUserInfo(ctx, pgCore, pgCoreQueries, userId)
	if getUserErr != nil {
		return userInfo, getUserErr
	}

	userInfo.Username = user.Username
	userInfo.UserHandle = user.UserHandle
	userInfo.IsAdult = user.IsAdult
	userInfo.ReadingHistory = user.ReadingHistory

	return userInfo, nil
}
