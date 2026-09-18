package pg_core_queries

import (
	"context"
	"freecreate/internal/config"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetUserInfo(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, userId int64) (pg_core_types.UserInfo, *api_error.Error) {
	var userInfo pg_core_types.UserInfo

	query := pgCoreQueries.GetUserInfo()

	namedArgs := pgx.NamedArgs{
		"user_id": userId,
	}

	err := pgCore.QueryRow(ctx, query, namedArgs).Scan(&userInfo)
	if err != nil {
		logger.Log(err)

		apiErr := &api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   err,
		}

		return userInfo, apiErr
	}

	return userInfo, nil
}
