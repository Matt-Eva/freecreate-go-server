package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UpdateUserInfoParams struct {
	Username       string
	UserHandle     string
	ReadingHistory bool
	IsAdult        bool
	UserId         int64
}

func HandleUpdateUserInfo(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params UpdateUserInfoParams) (UserInfo, *api_error.Error) {
	var userInfo UserInfo

	updateParams := pg_core_types.UpdateUserInfoParams{
		Username:       params.Username,
		UserHandle:     params.UserHandle,
		ReadingHistory: params.ReadingHistory,
		IsAdult:        params.IsAdult,
		UserId: params.UserId,
	}

	info, getInfoErr := pg_core_queries.UpdateUserInfo(ctx, pgCore, pgCoreQueries, updateParams)
	if getInfoErr != nil {
		return userInfo, getInfoErr
	}

	userInfo.Username = info.Username
	userInfo.UserHandle = strings.TrimLeft(info.UserHandle, "@")
	userInfo.ReadingHistory = info.ReadingHistory
	userInfo.IsAdult = info.IsAdult

	return userInfo, nil
}
