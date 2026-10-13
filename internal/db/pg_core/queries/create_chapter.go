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

func CreateChapter(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.CreateChapterParams) (pg_core_types.CreatedChapter, *api_error.Error){
	query := pgCoreQueries.CreateChapter()

	namedArgs := pgx.NamedArgs{
		"user_id": params.UserId,
		"creator_id": params.CreatorId,
		"writing_id": params.WritingId,
		"title": params.Title,
		"chapter_number": params.ChapterNumber,
		"content_db": 1,
	}

	var createdChapter pg_core_types.CreatedChapter

	queryErr := pgCore.QueryRow(ctx, query, namedArgs).Scan(&createdChapter)
	if queryErr != nil {
		logger.Log(queryErr)
		
		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: queryErr,
		}

		return createdChapter, apiErr
	}

	return createdChapter, nil
}