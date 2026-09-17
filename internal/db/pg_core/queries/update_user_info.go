package pg_core_queries

import (
	"context"
	"errors"
	"freecreate/internal/config"
	pg_core_types "freecreate/internal/db/pg_core/types"
	pg_core_validators "freecreate/internal/db/pg_core/validators"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

func UpdateUserInfo(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, updateUserInfoParams pg_core_types.UpdateUserInfoParams) (pg_core_types.UserInfo, *api_error.Error) {
	var userInfo pg_core_types.UserInfo

	namedArgs := pgx.NamedArgs{
		"username":        updateUserInfoParams.Username,
		"user_handle":     updateUserInfoParams.UserHandle,
		"reading_history": updateUserInfoParams.ReadingHistory,
		"is_adult":        updateUserInfoParams.IsAdult,
		"user_id":         updateUserInfoParams.UserId,
	}

	argErr := pg_core_validators.ValidateUserInfo(namedArgs)
	if argErr != nil {
		return userInfo, argErr
	}

	query := pgCoreQueries.UpdateUserInfo()

	var username string
	var user_handle string
	var is_adult bool
	var reading_history bool

	queryErr := pgCore.QueryRow(ctx, query, namedArgs).Scan(&username, &user_handle, &reading_history, &is_adult)

	var pgErr *pgconn.PgError
	if errors.As(queryErr, &pgErr) && pgErr.Code == "23505" {
		apiErr := &api_error.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "That user handle is already in use.",
			Error:   queryErr,
		}

		return userInfo, apiErr
	} else if queryErr != nil {
		logger.Log(queryErr)

		apiErr := &api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   queryErr,
		}

		return userInfo, apiErr
	}

	userInfo.Username = username
	userInfo.UserHandle = user_handle
	userInfo.IsAdult = is_adult
	userInfo.ReadingHistory = reading_history

	return userInfo, nil
}
