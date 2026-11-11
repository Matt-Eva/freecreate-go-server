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

func GetMyChapter(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.GetMyChapterParams)(pg_core_types.Chapter, *api_error.Error){
	var chapter pg_core_types.Chapter

	query := pgCoreQueries.GetMyChapter()

	args := pgx.NamedArgs{
		"user_id": params.UserId,
		"@chapter_uuid": params.UUID,
	}

	queryErr := pgCore.QueryRow(ctx, query, args).Scan(&chapter)
	if queryErr != nil {
		logger.Log(queryErr)
		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: queryErr,
		}

		return chapter, apiErr
	}

	return chapter, nil
}