package pg_core_queries

import (
	"context"
	"errors"
	"fmt"
	"freecreate/internal/config"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MyCreator struct {
	Name string
	UUID uuid.UUID
}

func GetMyCreators(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, userId int64) ([]MyCreator, *api_error.Error) {
	query := pgCoreQueries.GetMyCreators()
	namedArgs := pgx.NamedArgs{
		"user_id": userId,
	}

	queryResult, queryErr := pgCore.Query(ctx, query, namedArgs)
	if errors.Is(queryErr, pgx.ErrNoRows){
		return []MyCreator{}, nil
	} else if queryErr != nil {
		logger.Log(queryErr)

		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: queryErr,
		}
		return []MyCreator{}, apiErr
	}

	fmt.Println(queryResult)

	return []MyCreator{}, nil
}
