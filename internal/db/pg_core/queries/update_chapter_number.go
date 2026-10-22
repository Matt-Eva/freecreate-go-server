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

func UpdateChapterNumber(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.UpdateChapterNumberParams) (int, *api_error.Error) {

	query := pgCoreQueries.UpdateChapterNumber()

	namedArgs := pgx.NamedArgs{
		"user_id": params.UserId,
		"uuid": params.UUID,
		"chapter_number": params.ChapterNumber,
	}

	var newChapterNumber int

	queryErr := pgCore.QueryRow(ctx, query, namedArgs).Scan(&newChapterNumber)
	if queryErr != nil {
		logger.Log(queryErr)

		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: queryErr,
		}

		return newChapterNumber, apiErr
	}

	return newChapterNumber, nil
}
