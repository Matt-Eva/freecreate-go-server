package pg_content_queries

import (
	"context"
	"errors"
	"freecreate/internal/config"
	pg_content_types "freecreate/internal/db/pg_content/types"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)


func GetJsonChapterContent(ctx context.Context, pgContentPools config.PgContentPools, pgContentQueries config.PgContentQueries, params pg_content_types.GetJsonChapterContentParams)(pg_content_types.JsonChapterContent, *api_error.Error){
	var jsonChapterContent pg_content_types.JsonChapterContent

	var contentDb *pgxpool.Pool

	dbNumber := params.ContentDB

	switch dbNumber {
	case 1:
		contentDb = pgContentPools.PgContentOne
	default:
		err := errors.New("that is not a valid database number")
		logger.Log(err)

		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: err,
		}

		return jsonChapterContent, apiErr
	}

	query := pgContentQueries.GetJsonChapterContent()

	args := pgx.NamedArgs {
		"chapter_uuid": params.ChapterUuid,
	}

	queryErr := contentDb.QueryRow(ctx, query, args).Scan(&jsonChapterContent)
	if queryErr != nil {
		logger.Log(queryErr)
		
		apiErr := &api_error.Error{
			Code: http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error: queryErr,
		}

		return jsonChapterContent, apiErr
	}

	return jsonChapterContent, nil
}