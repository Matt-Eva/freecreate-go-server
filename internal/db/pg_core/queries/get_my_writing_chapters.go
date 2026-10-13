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

func GetMyWritingChapters(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.GetMyWritingChapterParams) ([]pg_core_types.Chapter, *api_error.Error) {
	var myWritingChapters []pg_core_types.Chapter

	query := pgCoreQueries.GetMyWritingChapters()

	namedArgs := pgx.NamedArgs{
		"user_id":    params.UserId,
		"writing_id": params.WritingId,
	}

	queryResult, queryErr := pgCore.Query(ctx, query, namedArgs)
	if queryErr != nil {
		logger.Log(queryErr)

		apiErr := &api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   queryErr,
		}

		return myWritingChapters, apiErr
	}

	for queryResult.Next() {
		var chapter pg_core_types.Chapter

		scanErr := queryResult.Scan(&chapter)
		if scanErr != nil {
			logger.Log(scanErr)

			apiErr := &api_error.Error{
				Code:    http.StatusInternalServerError,
				Message: api_error.InteralServerErrorMessage,
				Error:   scanErr,
			}

			return myWritingChapters, apiErr
		}

		myWritingChapters = append(myWritingChapters, chapter)
	}

	return myWritingChapters, nil
}
