package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateCreatorParams struct {
	UserId        int64
	Name          string
	CreatorHandle string
}

type CreatedCreator struct {
	Name          string
	CreatorHandle string
	UUID          uuid.UUID
}

func HandleCreateCreator(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params CreateCreatorParams) (CreatedCreator, *api_error.Error) {
	var createdCreator CreatedCreator

	queryParams := pg_core_types.CreateCreatorParams{
		UserId:        params.UserId,
		Name:          params.Name,
		CreatorHandle: params.CreatorHandle,
	}

	creator, createCreatorErr := pg_core_queries.CreateCreator(ctx, pgCore, pgCoreQueries, queryParams)
	if createCreatorErr != nil {
		return createdCreator, createCreatorErr
	}

	createdCreator = CreatedCreator{
		Name:          creator.Name,
		UUID:          creator.UUID,
		CreatorHandle: creator.CreatorHandle,
	}

	return createdCreator, nil
}
