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

func CreateChapterContent(ctx context.Context, dbNumber int16, pgContentPools config.PgContentPools, pgContentQueries config.PgContentQueries, params pg_content_types.CreateChapterContentParams) *api_error.Error {

	var contentDb *pgxpool.Pool

	switch dbNumber {
	case 1:
		contentDb = pgContentPools.PgContentOne
	default:
		err := errors.New("that is not a valid database number")
		logger.Log(err)
		apiErr := &api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   err,
		}
		return apiErr
	}

	query := pgContentQueries.CreateChapterContent()

	namedArgs := pgx.NamedArgs{
		"chapter_uuid": params.ChapterUUID,
	}

	_, err := contentDb.Exec(ctx, query, namedArgs)
	if err != nil {
		logger.Log(err)

		apiErr := &api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   err,
		}

		return apiErr
	}

	return nil
}
