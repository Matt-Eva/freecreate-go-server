package query_handlers

import (
	"context"
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	"freecreate/internal/lib/api_error"

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
	UUID string
}

func HandleCreateCreator(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries)(CreatedCreator, *api_error.Error){
	var createdCreator CreatedCreator

	createdCreator, createCreatorErr := pg_core_queries.CreateCreator(ctx, pgCore, pgCoreQueries, creatorName, userId)
	
	return createdCreator, nil
}