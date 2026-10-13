package pg_core_queries

import (
	"context"
	"freecreate/internal/config"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"net/http"
	"slices"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func UpdateWriting(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params pg_core_types.UpdateWritingParams) (pg_core_types.UpdatedWriting, *api_error.Error) {
	var updatedWriting pg_core_types.UpdatedWriting

	query := pgCoreQueries.UpdateWriting()

	sortedTopics := params.Topics
	slices.Sort(sortedTopics)

	var topics []string
	for _, val := range sortedTopics {
		if len(topics) >= 3 {
			break
		} else if len(sortedTopics) > 1 && (val == "No Topic" || val == "No Genre") {
			continue
		}

		topics = append(topics, val)
	}

	sortedTags := params.Tags
	slices.Sort(sortedTags)

	var tags []string
	for _, val := range params.Tags {
		if len(tags) >= 20 {
			break
		}
		tags = append(tags, val)
	}

	namedArgs := pgx.NamedArgs{
		"user_id":      params.UserId,
		"uuid":         params.UUID,
		"title":        params.Title,
		"subtitle":     params.Subtitle,
		"writing_type": params.WritingType,
		"description":  params.Description,
		"topics":       topics,
		"tags":         tags,
		"is_adult":     params.IsAdult,
		"creator_id":   params.CreatorId,
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
