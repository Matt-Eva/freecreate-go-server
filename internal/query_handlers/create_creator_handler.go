package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	"freecreate/internal/lib/api_error"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateCreatorParams struct {
	UserId int64
	Name string
	Handle string
}

type CreatedCreator struct {
	Name string
	Handle string
	UUID uuid.UUID
}

func HandleCreateCreator(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params CreateCreatorParams)(CreatedCreator, *api_error.Error){
	var createdCreator CreatedCreator

	queryParams := pg_core_queries.CreateCreatorParams{
		UserId: params.UserId,
		Name: params.Name,
		Handle: params.Handle,
	}

	creator, createCreatorErr := pg_core_queries.CreateCreator(ctx, pgCore, pgCoreQueries, queryParams)
	if createCreatorErr != nil {
		return createdCreator, createCreatorErr
	}

	createdCreator = CreatedCreator{
		Name: creator.Name,
		UUID: creator.UUID,
		Handle: creator.Handle,
	}
	
	return createdCreator, nil
}