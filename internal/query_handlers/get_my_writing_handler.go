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

type GetMyWritingParams struct {
	UserId int64
}

type MyWritingsStruct struct {
	Title string
	UUID uuid.UUID
	Published bool
	CreatorId int64
}

type CreatorWritingsGroup struct {
	CreatorName string
	CreatorUUID uuid.UUID
	Writing []MyWritingsStruct
}



func HandleGetMyWriting(ctx context.Context, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, params GetMyWritingParams)([]CreatorWritingsGroup, *api_error.Error) {
	var creatorWritings []CreatorWritingsGroup

	getWritingParams := pg_core_types.GetMyWritingsParams{
		UserId: params.UserId,
	}

	_, myWritingErr := pg_core_queries.GetMyWritings(ctx, pgCore, pgCoreQueries, getWritingParams)
	if myWritingErr != nil {
		return creatorWritings, myWritingErr
	}

	return creatorWritings, nil
}
