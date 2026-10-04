package pg_core_queries

import (
	"context"
	"freecreate/internal/config"
	pg_core_types "freecreate/internal/db/pg_core/types"
	pg_core_validators "freecreate/internal/db/pg_core/validators"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateWriting(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.NewWritingParams) (pg_core_types.CreatedWriting, *api_error.Error) {
	var createdWriting pg_core_types.CreatedWriting

	query := pgCoreQueries.CreateWriting()

	namedArgs := pgx.NamedArgs{
		"user_id":      params.UserId,
		"creator_id":   params.CreatorId,
		"title":        params.Title,
		"writing_type": params.WritingType,
	}

	argErr := pg_core_validators.ValidateNewWriting(namedArgs)
	if argErr != nil {
		return createdWriting, argErr
	}

	var uuid uuid.UUID
	var user_id int64
	var creator_id int64
	var title string
	var subtitle string
	var writing_type string
	var topics []string
	var tags []string
	var description string
	var published bool

	queryErr := pgCore.QueryRow(ctx, query, namedArgs).Scan(&uuid, &user_id, &creator_id, &title, &subtitle, &writing_type, &topics, &tags, &description, &published)
	if queryErr != nil {
		logger.Log(queryErr)
		apiErr := &api_error.Error{
			Code:    http.StatusInternalServerError,
			Message: api_error.InteralServerErrorMessage,
			Error:   queryErr,
		}
		return createdWriting, apiErr
	}

	createdWriting.UUID = uuid
	createdWriting.UserId = user_id
	createdWriting.CreatorId = creator_id
	createdWriting.Title = title
	createdWriting.Subtitle = subtitle
	createdWriting.Topics = topics
	createdWriting.Tags = tags
	createdWriting.Description = description
	createdWriting.Published = published

	return createdWriting, nil
}
