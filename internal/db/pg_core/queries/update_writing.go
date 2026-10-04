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

func UpdateWriting(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.UpdateWritingParams) (pg_core_types.UpdatedWriting, *api_error.Error) {
	var updatedWriting pg_core_types.UpdatedWriting

	query := pgCoreQueries.UpdateWriting()

	namedArgs := pgx.NamedArgs{
		"user_id":      params.UserId,
		"uuid":         params.UUID,
		"title":        params.Title,
		"subtitle":     params.Subtitle,
		"writing_type": params.WritingType,
		"description":  params.Description,
		"topics":       params.Topics,
		"tags":         params.Tags,
		"is_adult":     params.IsAdult,
	}

	queryErr := pgCore.QueryRow(ctx, query, namedArgs).Scan(&updatedWriting)
	if queryErr != nil {
		logger.Log(queryErr)
		apiErr := &api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   queryErr,
		}
		return updatedWriting, apiErr
	}

	return updatedWriting, nil
}
